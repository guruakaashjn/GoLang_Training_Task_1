package service

import (
	custom_errors "bankingapp/errors"
	"bankingapp/models/account"
	"bankingapp/models/bank"
	"bankingapp/models/bank_entry"
	"bankingapp/models/bank_passbook"
	"bankingapp/models/customer"
	"bankingapp/models/entry"
	"bankingapp/models/offer"
	"bankingapp/models/passbook"
	"bankingapp/repository"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type AccountService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewAccountService(db *gorm.DB, repo repository.Repository) *AccountService {
	return &AccountService{
		db:           db,
		repository:   repo,
		associations: []string{"Passbook", "Passbook.Entries", "Offers"},
	}

}

func (accountService *AccountService) doesAccountExist(Id uint) error {
	exists, err := repository.DoesRecordExist(accountService.db, int(Id), account.Account{}, repository.Filter("`id` = ?", Id))

	if !exists || err != nil {
		return errors.New("account id is invalid")
	}
	return nil
}

func (accountService *AccountService) CreateAccount(newAccount *account.Account) error {
	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	requiredCustomer := &customer.Customer{}
	err := accountService.repository.GetRecordForId(uow, newAccount.CustomerID, requiredCustomer)
	if err != nil {
		uow.RollBack()
		return custom_errors.NewValidationError("customer is deleted or does not exist")
	}

	requiredBank := &bank.Bank{}
	err = accountService.repository.GetRecordForId(uow, newAccount.BankID, requiredBank)
	if err != nil {
		uow.RollBack()
		return custom_errors.NewValidationError("bank is deleted or does not exist")
	}

	if newAccount.Balance < 1000 {
		uow.RollBack()
		return custom_errors.NewValidationError("balance is less then 1000")
	}

	err = accountService.repository.Add(uow, newAccount)
	if err != nil {
		uow.RollBack()
		return err
	}

	var passbook *passbook.Passbook = passbook.NewPassbook(newAccount.CustomerID, newAccount.ID, newAccount.Balance)
	passbook.AccountID = newAccount.ID
	err = accountService.repository.Add(uow, passbook)
	if err != nil {
		uow.RollBack()
		return err
	}

	var entry *entry.Entry = entry.NewEntry(newAccount.CustomerID, newAccount.CustomerID, newAccount.ID, newAccount.ID, newAccount.Balance, "CREDIT")
	entry.PassbookID = passbook.ID
	err = accountService.repository.Add(uow, entry)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}
func (accountService *AccountService) GetAllAccounts(allAccounts *[]account.Account, totalCount *int, customerId uint, limit int, offset int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(accountService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(accountService.associations, givenAssociations)

	err := accountService.repository.GetAll(uow, allAccounts, repository.Filter("`customer_id` = ?", customerId), repository.Paginate(limit, offset, totalCount), repository.Preload(requiredAssociations))
	if err != nil {
		return err

	}
	uow.Commit()
	return nil
}

func (accountService *AccountService) GetAccountById(requiredAccount *account.Account, givenAssociations []string) error {
	err := accountService.doesAccountExist(requiredAccount.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(accountService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(accountService.associations, givenAssociations)
	err = accountService.repository.GetRecordForId(uow, requiredAccount.ID, requiredAccount, repository.Filter("`customer_id` =?", requiredAccount.CustomerID), repository.Preload(requiredAssociations))
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (accountService *AccountService) DeleteAccount(accountToDelete *account.Account) error {
	err := accountService.doesAccountExist(accountToDelete.ID)

	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(accountService.db, false)

	defer uow.RollBack()
	if err := accountService.repository.UpdateWithMap(uow, accountToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
		"IsActive":  false,
	},
		repository.Filter("`id`=?", accountToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (accountService *AccountService) UpdateAccount(accountToUpdate *account.Account) error {
	err := accountService.doesAccountExist(accountToUpdate.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()
	tempAccount := account.Account{}

	err = accountService.repository.GetRecordForId(uow, accountToUpdate.ID, &tempAccount, repository.Select("`created_at`"), repository.Filter("`id` = ?", accountToUpdate.ID))
	if err != nil {
		return err
	}
	accountToUpdate.CreatedAt = tempAccount.CreatedAt
	err = accountService.repository.Save(uow, accountToUpdate)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (accountService *AccountService) TransferMoney(senderAccount, receiverAccount *account.Account, amount int) error {
	err := accountService.doesAccountExist(senderAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("sender account is deleted or does not exist")
	}
	err = accountService.doesAccountExist(receiverAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("receiver account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	senderAccountTemp := &account.Account{}
	receiverAccountTemp := &account.Account{}
	err = accountService.repository.GetRecord(uow, senderAccountTemp, repository.Filter("`id` =? ", senderAccount.ID))
	if err != nil {
		return err
	}
	err = accountService.repository.GetRecord(uow, receiverAccountTemp, repository.Filter("`id` =? ", receiverAccount.ID))
	if err != nil {
		return err
	}

	if senderAccountTemp.Balance < amount {
		return custom_errors.NewValidationError("sender account in-sufficient balance")
	}
	senderAccountTemp.Balance -= amount
	receiverAccountTemp.Balance += amount

	err = accountService.repository.Save(uow, senderAccountTemp)
	if err != nil {
		return err

	}
	err = accountService.repository.Save(uow, receiverAccountTemp)
	if err != nil {
		return err
	}
	senderAccountPassbook := &passbook.Passbook{}
	receiverAccountPassbook := &passbook.Passbook{}

	err = accountService.repository.GetRecord(uow, senderAccountPassbook, repository.Filter("`account_id` = ?", senderAccountTemp.ID))
	if err != nil {
		return err
	}

	err = accountService.repository.GetRecord(uow, receiverAccountPassbook, repository.Filter("`account_id` = ?", receiverAccountTemp.ID))
	if err != nil {
		return err
	}

	senderEntry := entry.NewEntry(
		senderAccountTemp.CustomerID,
		receiverAccountTemp.CustomerID,
		senderAccountTemp.ID,
		receiverAccountTemp.ID,
		amount,
		"DEBIT",
	)
	senderEntry.PassbookID = senderAccountPassbook.ID

	receiverEntry := entry.NewEntry(
		senderAccountTemp.CustomerID,
		receiverAccountTemp.CustomerID,
		senderAccountTemp.ID,
		receiverAccountTemp.ID,
		amount,
		"CREDIT",
	)
	receiverEntry.PassbookID = receiverAccountPassbook.ID

	err = accountService.repository.Add(uow, senderEntry)
	if err != nil {
		return err
	}

	err = accountService.repository.Add(uow, receiverEntry)
	if err != nil {
		return err

	}

	senderBank := &bank.Bank{}
	receiverBank := &bank.Bank{}

	err = accountService.repository.GetRecord(uow, senderBank, repository.Filter("`id` = ? ", senderAccountTemp.BankID))
	if err != nil {
		return err
	}

	err = accountService.repository.GetRecord(uow, receiverBank, repository.Filter("`id` = ?", receiverAccountTemp.BankID))
	if err != nil {
		return err
	}

	senderBankPassbook := &bank_passbook.BankPassbook{}
	receiverBankPassbook := &bank_passbook.BankPassbook{}

	err = accountService.repository.GetRecord(uow, senderBankPassbook, repository.Filter("`bank_id` = ?", senderBank.ID))
	if err != nil {
		return err
	}
	err = accountService.repository.GetRecord(uow, receiverBankPassbook, repository.Filter("`bank_id` = ?", receiverBank.ID))
	if err != nil {
		return err
	}

	senderBankEntry := bank_entry.NewBankEntry(
		senderBank.ID,
		receiverBank.ID,
		amount,
		"DEBIT",
	)
	senderBankEntry.BankPassbookID = senderBankPassbook.ID

	receiverBankEntry := bank_entry.NewBankEntry(
		senderBank.ID,
		receiverBank.ID,
		amount,
		"CREDIT",
	)
	receiverBankEntry.BankPassbookID = receiverBankPassbook.ID

	err = accountService.repository.Add(uow, senderBankEntry)
	if err != nil {
		return err

	}
	err = accountService.repository.Add(uow, receiverBankEntry)
	if err != nil {
		return err

	}

	uow.Commit()
	return nil

}

func (accountService *AccountService) DepositMoney(senderAccount *account.Account, amount int) error {
	err := accountService.doesAccountExist(senderAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("sender account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	senderAccountTemp := account.Account{}
	err = accountService.repository.GetRecord(uow, &senderAccountTemp, repository.Filter("`id` =? ", senderAccount.ID))
	if err != nil {
		return err
	}

	senderAccountTemp.Balance += amount

	err = accountService.repository.Save(uow, senderAccountTemp)
	if err != nil {
		return err

	}

	senderEntry := entry.NewEntry(
		senderAccountTemp.CustomerID,
		senderAccountTemp.CustomerID,
		senderAccountTemp.ID,
		senderAccountTemp.ID,
		amount,
		"CREDIT",
	)

	err = accountService.repository.Add(uow, senderEntry)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil

}

func (accountService *AccountService) WithdrawMoney(senderAccount *account.Account, amount int) error {
	err := accountService.doesAccountExist(senderAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("sender account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	senderAccountTemp := account.Account{}
	err = accountService.repository.GetRecord(uow, &senderAccountTemp, repository.Filter("`id` =? ", senderAccount.ID))
	if err != nil {
		return err

	}

	if senderAccountTemp.Balance < amount {
		return custom_errors.NewValidationError("sender account in-sufficient balance")
	}

	senderAccountTemp.Balance -= amount

	err = accountService.repository.Save(uow, senderAccountTemp)
	if err != nil {
		return err

	}

	senderEntry := entry.NewEntry(
		senderAccountTemp.CustomerID,
		senderAccountTemp.CustomerID,
		senderAccountTemp.ID,
		senderAccountTemp.ID,
		amount,
		"DEBIT",
	)

	err = accountService.repository.Add(uow, senderEntry)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil

}

func (accountService *AccountService) PassbookPrint(requiredAccount *account.Account, startDateTimeDotTime time.Time, endDateTimeDotTime time.Time, requiredEntries *[]entry.Entry, limit int, offset int, totalCount *int) error {
	err := accountService.doesAccountExist(requiredAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	requiredAccountTemp := &account.Account{}
	err = accountService.repository.GetRecordForId(uow, requiredAccount.ID, requiredAccountTemp)
	if err != nil {
		return err

	}
	requiredPassbook := &passbook.Passbook{}
	err = accountService.repository.GetRecord(uow, requiredPassbook, repository.Filter("`account_id` = ?", requiredAccountTemp.ID))
	if err != nil {
		return err

	}
	err = accountService.repository.GetAll(uow, requiredEntries,
		repository.Paginate(limit, offset, totalCount),
		repository.Filter("`passbook_id` =?", requiredPassbook.ID),
		repository.Filter("`created_at` >= ? AND `created_at` <= ?", startDateTimeDotTime, endDateTimeDotTime),
	)
	if err != nil {
		return err

	}

	// fmt.Println(requiredEntries)
	uow.Commit()
	return nil

}

func (accountService *AccountService) GetAvailableOffers(requiredAccount *account.Account, availableOffers *[]offer.Offer, givenAssociations []string) error {
	err := accountService.doesAccountExist(requiredAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	err = accountService.repository.GetRecord(uow, requiredAccount,
		repository.Filter("`customer_id` =?", requiredAccount.CustomerID),
		repository.Filter("`id` =?", requiredAccount.ID),
	)

	if err != nil {
		return nil
	}

	err = accountService.repository.GetAll(uow, availableOffers,
		repository.Filter("`bank_id` =?", requiredAccount.BankID),
	)

	if err != nil {
		return nil
	}

	uow.Commit()
	return nil

}

func (accountService *AccountService) TakeAvailableOffer(requiredAccount *account.Account, takeOffer *offer.Offer) error {
	err := accountService.doesAccountExist(requiredAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	err = accountService.repository.GetRecord(uow, requiredAccount,
		repository.Filter("`customer_id` =?", requiredAccount.CustomerID),
		repository.Filter("`id` =?", requiredAccount.ID),
	)

	if err != nil {
		return nil
	}

	err = accountService.repository.GetRecordForId(uow, takeOffer.ID, takeOffer)
	if err != nil {
		return nil
	}

	if takeOffer.BankID != requiredAccount.BankID {
		return errors.New("please provide proper offer id of your bank only")
	}

	requiredAccount.Offers = append(requiredAccount.Offers, *takeOffer)
	requiredAccount.UpdatedAt = time.Now()
	err = accountService.repository.Update(uow, requiredAccount)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}

func (accountService *AccountService) ChosenOffers(requiredAccount *account.Account, chosenOffers *[]offer.Offer) error {
	err := accountService.doesAccountExist(requiredAccount.ID)
	if err != nil {
		return custom_errors.NewValidationError("account is deleted or does not exist")
	}

	uow := repository.NewUnitOfWork(accountService.db, false)
	defer uow.RollBack()

	err = accountService.repository.GetRecord(uow, requiredAccount,
		repository.Filter("`customer_id` =?", requiredAccount.CustomerID),
		repository.Filter("`id` =?", requiredAccount.ID),
	)

	if err != nil {
		return nil
	}

	err = accountService.repository.GetAll(uow, chosenOffers,
		repository.Table("offers"),
		repository.Filter("accounts.id=?", requiredAccount.ID),
		repository.Join("join account_offers_join on account_offers_join.offer_id = offers.id"),
		repository.Join("join accounts on account_offers_join.account_id = accounts.id"),
	)

	// fmt.Println(chosenOffers)
	if err != nil {
		return nil
	}

	uow.Commit()
	return nil
}

// package service

// import (
// 	"bankingapp/components/guru_passbook"
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// )

// type Account struct {
// 	AccountNumber uuid.UUID
// 	BankId        uuid.UUID
// 	CustomerId    uuid.UUID
// 	IsActive      bool
// 	Balance       int
// 	Passbook      *guru_passbook.Passbook
// }

// func NewAccount(BankId uuid.UUID, CustomerId uuid.UUID, Balance int) *Account {
// 	AccountNumber := uuid.New()
// 	var newPassbook *guru_passbook.Passbook = guru_passbook.NewPassbook(CustomerId, AccountNumber, Balance)
// 	fmt.Println("Passbook from NewAccount() guru_account.go service : ", newPassbook)
// 	var newAccountObject *Account = &Account{
// 		AccountNumber: AccountNumber,
// 		BankId:        BankId,
// 		CustomerId:    CustomerId,
// 		IsActive:      true,
// 		Balance:       Balance,
// 		Passbook:      newPassbook,
// 	}
// 	fmt.Println(newAccountObject.Passbook)

// 	return newAccountObject
// }

// func CreateAccount(BankId uuid.UUID, CustomerId uuid.UUID, Balance int) (account *Account) {

// 	return NewAccount(BankId, CustomerId, Balance)

// }

// func (a *Account) ReadAccount() (bool, *Account) {
// 	if a.IsActive {
// 		return true, a
// 	}
// 	return false, a

// }
// func (a *Account) UpdateAccount(updateField string, updateValue interface{}) *Account {
// 	// Balance, abbr

// 	switch updateValue := updateValue.(type) {
// 	case int:
// 		a.SetBalance(updateValue)
// 	case uuid.UUID:
// 		a.BankId = updateValue
// 	}

// 	return a

// }
// func (a *Account) DeleteAccount() *Account {

// 	a.SetIsActive()
// 	return a
// }

// func (a *Account) GetAccountNumber() uuid.UUID {
// 	return a.AccountNumber
// }
// func (a *Account) GetBankId() uuid.UUID {
// 	return a.BankId
// }
// func (a *Account) GetCustomerId() uuid.UUID {
// 	return a.CustomerId
// }
// func (a *Account) GetIsActive() bool {
// 	return a.IsActive
// }
// func (a *Account) SetIsActive() {
// 	a.IsActive = false
// }
// func (a *Account) GetBalance() int {
// 	return a.Balance
// }
// func (a *Account) SetBalance(Balance int) {
// 	a.Balance = Balance
// }

// func (a *Account) GetPassbook(startDate, endDate time.Time) *guru_passbook.Passbook {
// 	// fmt.Println("Get Passbook Func : ", startDate, "    ", endDate)
// 	return a.Passbook.ReadPassbook(startDate, endDate)
// }
// func (a *Account) DepositMoney(amount int) *Account {
// 	a.Balance += amount
// 	a.Passbook.AddEntry(a.CustomerId, a.CustomerId, a.AccountNumber, a.AccountNumber, amount, "CREDIT")
// 	return a
// }

// func (a *Account) WithdrawMoney(amount int) *Account {
// 	a.Balance -= amount
// 	a.Passbook.AddEntry(a.CustomerId, a.CustomerId, a.AccountNumber, a.AccountNumber, amount, "DEBIT")
// 	return a

// }

// func (a *Account) TransferMoney(receiver *Account, amount int) (*Account, *Account) {

// 	a.Balance -= amount
// 	receiver.Balance += amount
// 	a.Passbook.AddEntry(a.CustomerId, receiver.CustomerId, a.AccountNumber, receiver.AccountNumber, amount, "DEBIT")
// 	receiver.Passbook.AddEntry(a.CustomerId, receiver.CustomerId, a.AccountNumber, receiver.AccountNumber, amount, "CREDIT")
// 	return a, receiver
// }
