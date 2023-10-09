package service

import (
	"bankingapp/models/account"
	"bankingapp/models/bank"
	"bankingapp/models/bank_entry"
	"bankingapp/models/bank_passbook"
	"bankingapp/repository"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type BankService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewBankService(db *gorm.DB, repo repository.Repository) *BankService {
	return &BankService{
		db:           db,
		repository:   repo,
		associations: []string{"Accounts", "Accounts.Passbook", "Accounts.Passbook.Entries", "Accounts.Offers", "BankPassbook", "BankPassbook.Entries", "Offers"},
	}
}

func (bankService *BankService) doesBankExist(Id uint) error {
	exists, err := repository.DoesRecordExist(bankService.db, int(Id), bank.Bank{}, repository.Filter("`id` = ?", Id))

	if !exists || err != nil {
		return errors.New("data id is invalid")
	}
	return nil

}

func (bankService *BankService) setAbbreviation(uow *repository.UnitOfWork, FullName string) (abbr string) {
	abbr = FullName[0:2] + FullName[len(FullName)-2:]
	max := rand.Intn(10000-5000) + 5000
	min := rand.Intn(4999-0) + 0
	rnd := rand.Intn(max-min) + min
	abbr += strconv.Itoa(rnd)

	return abbr
}

func (bankService *BankService) CreateBank(newBank *bank.Bank) error {
	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()
label:
	tempAbbreviation := bankService.setAbbreviation(uow, newBank.FullName)
	err := bankService.repository.GetRecord(uow, bank.Bank{}, repository.Select("`abbreviation`"), repository.Filter("`abbreviation` = ?", tempAbbreviation))
	if err == nil {
		goto label
	}

	newBank.Abbreviation = tempAbbreviation
	newBank.IsActive = true
	err = bankService.repository.Add(uow, newBank)
	if err != nil {
		uow.RollBack()
		return err
	}

	var bank_passbook *bank_passbook.BankPassbook = bank_passbook.NewBankPassbook()
	bank_passbook.BankID = newBank.ID
	err = bankService.repository.Add(uow, bank_passbook)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}

func (bankService *BankService) GetAllBanks(allBanks *[]bank.Bank, totalCount *int, limit, offset int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(bankService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(bankService.associations, givenAssociations)

	err := bankService.repository.GetAll(uow, allBanks, repository.Paginate(limit, offset, totalCount), repository.Preload(requiredAssociations))

	if err != nil {
		return err
	}

	// *totalCount = len(*allBanks)
	uow.Commit()
	return nil
}

func (bankService *BankService) GetBankById(requiredBank *bank.Bank, idTemp int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(bankService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(bankService.associations, givenAssociations)

	err := bankService.repository.GetRecordForId(uow, uint(idTemp), requiredBank, repository.Preload(requiredAssociations))
	if err != nil {
		return err

	}
	uow.Commit()
	return nil

}

func (bankService *BankService) UpdateBank(bankToUpdate *bank.Bank) error {
	err := bankService.doesBankExist(bankToUpdate.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()

	tempBank := bank.Bank{}

	err = bankService.repository.GetRecordForId(uow, bankToUpdate.ID, &tempBank, repository.Filter("`id` = ?", bankToUpdate.ID))
	// fmt.Println("-------------------------------------", tempBank)
	if err != nil {
		return err
	}
	bankToUpdate.CreatedAt = tempBank.CreatedAt
	bankToUpdate.Abbreviation = tempBank.Abbreviation

	if bankToUpdate.FullName != tempBank.FullName {
	label:
		tempAbbreviation := bankService.setAbbreviation(uow, bankToUpdate.FullName)
		err = bankService.repository.GetRecord(uow, bank.Bank{}, repository.Select("`abbreviation`"), repository.Filter("`abbreviation` = ?", tempAbbreviation))
		if err == nil {
			goto label
		}
		bankToUpdate.Abbreviation = tempAbbreviation
	}

	err = bankService.repository.Save(uow, bankToUpdate)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (bankService *BankService) DeleteBank(bankToDelete *bank.Bank) error {
	err := bankService.doesBankExist(bankToDelete.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()
	fmt.Println(bankToDelete.ID)
	tempAccount := &account.Account{}
	check := bankService.repository.GetRecord(uow, tempAccount, repository.Filter("`bank_id` = ?", bankToDelete.ID))
	// fmt.Println("--------------------------------------------------", tempAccount)
	// fmt.Println("-------------------------------------------------", check)
	if check == nil {
		// fmt.Println("----------------------------------okay--------------------------------------------")
		return errors.New("bank cannot be deleted as some accounts exist")
	}
	if err := bankService.repository.UpdateWithMap(uow, bankToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
		"IsActive":  false,
	},
		repository.Filter("`id`=?", bankToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

type Trial struct {
	bank.Bank
	account.Account
}

func (bankService *BankService) AllBankNetWorth(mapAllBankNetworth map[string]int) error {
	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()

	trials := &[]Trial{}

	err := bankService.repository.GetAll(uow, trials,
		repository.Table("accounts"),
		repository.Join("join banks on accounts.bank_id = banks.id"),
		repository.Select("SUM(accounts.balance) AS `balance`, banks.full_name AS `full_name`"),
		repository.GroupBy("bank_id"),
	)

	if err != nil {
		return err
	}

	for _, trial := range *trials {
		mapAllBankNetworth[trial.FullName] = trial.Balance
	}

	uow.Commit()
	return nil
}

func (bankService *BankService) BankNetWorth(requiredBank *bank.Bank, mapBankNetWorth map[string]int) error {
	err := bankService.doesBankExist(requiredBank.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()

	trial := &Trial{}
	err = bankService.repository.GetAll(uow, trial,
		repository.Table("accounts"),
		repository.Filter("`bank_id` = ?", requiredBank.ID),
		repository.Join("join banks on accounts.bank_id = banks.id"),
		repository.Select("SUM(accounts.balance) AS `balance`, banks.full_name AS `full_name`"),
		repository.GroupBy("`bank_id`"),
	)
	if err != nil {
		return err
	}

	mapBankNetWorth[(*trial).FullName] = (*trial).Balance

	uow.Commit()
	return nil
}

type BankBankPassbookBankEntryJoin struct {
	bank.Bank
	bank_passbook.BankPassbook
	bank_entry.BankEntry
}

func (bankService *BankService) BankBalanceMap(requiredBank *bank.Bank, mapBankBalance map[string]map[string]int, startDateTimeDotTime, endDateTimeDotTime time.Time) error {
	err := bankService.doesBankExist(requiredBank.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()

	entries := &[]BankBankPassbookBankEntryJoin{}
	err = bankService.repository.GetAll(uow, entries,
		repository.Table("banks"),
		repository.Filter("banks.id = ?", requiredBank.ID),

		repository.Join("join bank_passbooks on banks.id = bank_passbooks.bank_id"),
		repository.Join("join bank_entries on bank_entries.bank_passbook_id = bank_passbooks.id"),
		repository.Select("banks.id AS `id`, banks.full_name AS `full_name`,bank_entries.amount AS `amount`, bank_entries.from_bank AS `from_bank`, bank_entries.to_bank AS `to_bank`, bank_entries.created_at AS `created_at`, bank_entries.transaction_type AS `transaction_type`"),
		repository.Filter("bank_entries.created_at >= ? AND bank_entries.created_at <= ?", startDateTimeDotTime, endDateTimeDotTime),
	)
	if err != nil {
		return err
	}
	// fmt.Println((*entries))
	mapBankBalanceTemp := make(map[uint]int, 0)
	bankArrayTemp := []bank.Bank{}
	bankTemp := bank.Bank{}
	for _, entry := range *entries {
		if entry.TransactionType == "CREDIT" {
			// mapBankBalanceTemp[entry.ToBank] -= entry.Amount
			mapBankBalanceTemp[entry.FromBank] += -entry.Amount
			bankTemp.ID = entry.FromBank
			bankArrayTemp = append(bankArrayTemp, bankTemp)
		} else {
			mapBankBalanceTemp[entry.ToBank] += entry.Amount
			// mapBankBalanceTemp[entry.FromBank] -= entry.Amount
			bankTemp.ID = entry.ToBank
			bankArrayTemp = append(bankArrayTemp, bankTemp)
		}
	}

	err = bankService.repository.GetAll(uow, &bankArrayTemp, repository.Select("`id`, `full_name`"))
	if err != nil {
		return err
	}
	// fmt.Println("-------------------------------", bankArrayTemp)

	mapBankBalanceName := make(map[string]int, 0)
	bankName := ""
	for _, value := range bankArrayTemp {
		yes, err := mapBankBalanceTemp[value.ID]
		if err {
			mapBankBalanceName[value.FullName] = yes
		}
		if value.ID == requiredBank.ID {
			bankName = value.FullName
		}

	}
	mapBankBalance[bankName] = mapBankBalanceName
	// fmt.Println(mapBankBalanceName)

	// fmt.Println(mapBankBalanceTemp)
	// fmt.Println(mapBankBalance)

	uow.Commit()
	return nil
}

func (bankService *BankService) AllBankBalanceMap(allMapBankBalance *[]map[string]map[string]int, startDateTimeDotTime, endDateTimeDotTime time.Time) error {

	uow := repository.NewUnitOfWork(bankService.db, false)
	defer uow.RollBack()

	banks := &[]bank.Bank{}

	err := bankService.repository.GetAll(uow, banks, repository.Select("`id`, `full_name`"))
	if err != nil {
		return err
	}

	for _, bankTemp := range *banks {
		requiredBank := &bank.Bank{}
		err = bankService.repository.GetRecordForId(uow, bankTemp.ID, requiredBank)
		if err != nil {
			return err
		}
		mapBankBalance := make(map[string]map[string]int)
		bankService.BankBalanceMap(requiredBank, mapBankBalance, startDateTimeDotTime, endDateTimeDotTime)
		*allMapBankBalance = append(*allMapBankBalance, mapBankBalance)
	}

	// fmt.Println(allMapBankBalance)

	uow.Commit()
	return nil
}

// package guru_bank

// import (
// 	account_service "bankingapp/components/guru_account/service"
// 	"bankingapp/components/guru_bank_passbook"
// 	"math/rand"
// 	"strconv"
// 	"time"

// 	"github.com/google/uuid"
// )

// var Banks = make([]*Bank, 0)

// type Bank struct {
// 	BankId       uuid.UUID
// 	FullName     string
// 	Abbreviation string
// 	IsActive     bool
// 	Accounts     []*account_service.Account
// 	BankPassbook *guru_bank_passbook.BankPassbook
// }

// func NewBank(FullName string) *Bank {
// 	var abbr string = setAbbreviation(FullName)
// 	var initialAccountsList []*account_service.Account = make([]*account_service.Account, 0)

// 	var bankPassbookInitial *guru_bank_passbook.BankPassbook = guru_bank_passbook.CreateBankPassbook()
// 	var newBankObject = &Bank{
// 		BankId:       uuid.New(),
// 		FullName:     FullName,
// 		Abbreviation: abbr,
// 		IsActive:     true,
// 		Accounts:     initialAccountsList,

// 		BankPassbook: bankPassbookInitial,
// 	}
// 	Banks = append(Banks, newBankObject)
// 	return newBankObject
// }

// func (b *Bank) GetAbbreviation() string {
// 	return b.Abbreviation
// }

// func setAbbreviation(FullName string) (abbr string) {
// 	abbr = FullName[0:4] + FullName[len(FullName)-4:]
// 	max := rand.Intn(10000-5000) + 5000
// 	min := rand.Intn(4999-0) + 0
// 	rnd := rand.Intn(max-min) + min
// 	abbr += strconv.Itoa(rnd)

// 	if checkAbbreviation(abbr) {
// 		return abbr
// 	}

// 	return setAbbreviation(FullName)
// }

// func checkAbbreviation(abbr string) (flag bool) {
// 	for i := 0; i < len(Banks); i++ {
// 		if Banks[i].GetAbbreviation() == abbr {
// 			return false
// 		}

// 	}
// 	return true
// }

// func CreateBank(FullName string) (bank *Bank) {

// 	return NewBank(FullName)

// }

// func (b *Bank) ReadBank() (bool, *Bank) {
// 	if b.IsActive {
// 		return true, b
// 	}
// 	return false, b

// }
// func ReadBankById(bankIdTemp uuid.UUID) (bool, *Bank) {
// 	var bank *Bank
// 	for i := 0; i < len(Banks); i++ {
// 		if Banks[i].BankId == bankIdTemp {
// 			bank = Banks[i]
// 			break
// 		}
// 	}
// 	if bank.IsActive {
// 		return true, bank
// 	}
// 	return false, bank
// }
// func ReadAllBanks() []*Bank {

// 	var allBanks []*Bank
// 	for i := 0; i < len(Banks); i++ {
// 		if Banks[i].IsActive {
// 			allBanks = append(allBanks, Banks[i])
// 		}
// 	}

// 	return allBanks

// }

// func (b *Bank) UpdateBank(updateValue string) *Bank {

// 	b.FullName = updateValue
// 	b.Abbreviation = setAbbreviation(b.FullName)
// 	return b

// }

// func (b *Bank) UpdateBankObject(bankTempObject *Bank) *Bank {
// 	if bankTempObject.FullName != "" && bankTempObject.FullName != b.FullName {
// 		b.FullName = bankTempObject.FullName
// 		b.Abbreviation = setAbbreviation(b.FullName)
// 	}

// 	return b

// }

// func (b *Bank) DeleteBank() *Bank {

// 	b.IsActive = false
// 	return b

// }

// // HELPER FUNCTIONS
// func (b *Bank) GetBankId() uuid.UUID {
// 	return b.BankId
// }
// func (b *Bank) GetBankName() string {
// 	return b.FullName
// }
// func (b *Bank) GetNetWorthOfBank() (networth int) {

// 	for i := 0; i < len(b.Accounts); i++ {
// 		networth += b.Accounts[i].GetBalance()
// 	}

// 	return networth
// }

// func (b *Bank) CheckBankContainsActiveAccounts() bool {
// 	var flag bool = false
// 	for i := 0; i < len(b.Accounts); i++ {
// 		if b.Accounts[i].GetIsActive() {
// 			flag = true
// 			break
// 		}
// 	}
// 	return flag
// }

// func (b *Bank) GetBankPassbook() *guru_bank_passbook.BankPassbook {
// 	return b.BankPassbook
// }

// func (b *Bank) ReadPassbookFromRange(fromDate time.Time, toDate time.Time) map[uuid.UUID]int {

// 	return b.BankPassbook.ReadAllEntries(fromDate, toDate)
// }
