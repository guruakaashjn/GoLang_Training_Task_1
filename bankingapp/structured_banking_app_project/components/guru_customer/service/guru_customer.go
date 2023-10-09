package service

import (
	"bankingapp/models/account"
	"bankingapp/models/customer"
	"bankingapp/repository"
	"bankingapp/utils"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewCustomerService(db *gorm.DB, repo repository.Repository) *CustomerService {
	return &CustomerService{
		db:           db,
		repository:   repo,
		associations: []string{"Accounts", "Accounts.Passbook", "Accounts.Passbook.Entries", "Accounts.Offers"},
	}
}

func (customerService *CustomerService) doesCustomerExist(Id uint) error {
	exists, err := repository.DoesRecordExist(customerService.db, int(Id), customer.Customer{}, repository.Filter("`id` = ?", Id))

	if !exists || err != nil {
		return errors.New("data id is invalid")
	}
	return nil
}

func (customerService *CustomerService) CreateCustomer(newCustomer *customer.Customer) error {
	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()

	HashedPassword, _ := utils.GenerateHash(newCustomer.Password)
	newCustomer.Password = string(HashedPassword)

	err := customerService.repository.Add(uow, newCustomer)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (customerService *CustomerService) GetAllCustomers(allCustomers *[]customer.Customer, totalCount *int, limit, offset int, givenAssociations []string, columnNames, conditions, operators, values []string) error {
	uow := repository.NewUnitOfWork(customerService.db, true)

	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(customerService.associations, givenAssociations)

	// fmt.Println(requiredAssociations)
	valuesRough := make([]interface{}, 0)
	valuesRough = append(valuesRough, values)
	err := customerService.repository.GetAll(uow, allCustomers,
		repository.Table("customers"),
		repository.Paginate(limit, offset, totalCount),
		repository.Preload(requiredAssociations),
		repository.FilterWithOperator(columnNames, conditions, operators, valuesRough),
	)

	if err != nil {
		return err
	}

	// *totalCount = len(*allCustomers)
	uow.Commit()
	return nil
}

func (customerService *CustomerService) GetCustomerById(requiredCustomer *customer.Customer, idTemp int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(customerService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(customerService.associations, givenAssociations)
	err := customerService.repository.GetRecordForId(uow, uint(idTemp), requiredCustomer, repository.Preload(requiredAssociations))

	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (customerService *CustomerService) UpdateCustomer(customerToUpdate *customer.Customer) error {
	err := customerService.doesCustomerExist(customerToUpdate.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()

	tempCustomer := customer.Customer{}

	err = customerService.repository.GetRecordForId(uow, customerToUpdate.ID, &tempCustomer, repository.Select("`created_at`"), repository.Filter("`id` = ?", customerToUpdate.ID))

	if err != nil {
		return err

	}

	customerToUpdate.CreatedAt = tempCustomer.CreatedAt
	err = customerService.repository.Save(uow, customerToUpdate)
	if err != nil {
		return err

	}

	uow.Commit()
	return nil
}

func (customerService *CustomerService) DeleteCustomer(customerToDelete *customer.Customer) error {
	err := customerService.doesCustomerExist(customerToDelete.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()

	if err := customerService.repository.UpdateWithMap(uow, customerToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
		"IsActive":  false,
	}, repository.Filter("`id`=?", customerToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (customerService *CustomerService) TotalBalance(currentCustomer *customer.Customer, totalBalanceAnswer *int) error {
	err := customerService.doesCustomerExist(currentCustomer.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()

	allAccountsBalance := &[]account.Account{}

	err = customerService.repository.GetAll(uow, allAccountsBalance,
		repository.Select("SUM(balance) AS `balance`"),
		repository.Filter("`customer_id` = ?", currentCustomer.ID),
	)
	if err != nil {
		return nil
	}

	// currentCustomerTemp := &customer.Customer{}
	// err = customerService.repository.GetRecordForId(uow, currentCustomer.ID, currentCustomerTemp, []string{})
	// if err != nil {
	// 	return err
	// }

	// var totalBalance int = 0
	// for _, account := range *allAccountsBalance {
	// 	totalBalance += account.Balance
	// }
	// fmt.Println(totalBalance)

	// currentCustomerTemp.TotalBalance = totalBalance
	// err = customerService.repository.Save(uow, currentCustomerTemp)
	// if err != nil {
	// 	return err

	// }
	// currentCustomer.TotalBalance = currentCustomerTemp.TotalBalance
	// _ = totalBalanceAnswer
	*totalBalanceAnswer = (*allAccountsBalance)[0].Balance
	// fmt.Println(*totalBalanceAnswer)
	uow.Commit()
	return nil
}

func (customerService *CustomerService) AccountBalanceList(currentCustomer *customer.Customer, mapAccountBalance map[uint]int) error {
	err := customerService.doesCustomerExist(currentCustomer.ID)
	if err != nil {
		return err
	}
	// fmt.Println(mapAccountBalance)
	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()

	allAccountsBalance := &[]account.Account{}

	err = customerService.repository.GetAll(uow, allAccountsBalance, repository.Filter("`customer_id` = ?", currentCustomer.ID), repository.Select("`balance`, `id`"))
	if err != nil {
		return nil
	}
	// fmt.Println(allAccountsBalance)

	for _, account := range *allAccountsBalance {
		mapAccountBalance[account.ID] = account.Balance
	}
	// fmt.Println(mapAccountBalance)
	uow.Commit()
	return nil
}

// package guru_customer

// import (
// 	account_service "bankingapp/components/guru_account/service"
// 	bank_service "bankingapp/components/guru_bank/service"
// 	"bankingapp/components/guru_passbook"
// 	"bankingapp/guru_errors"
// 	"bankingapp/utils"
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// )

// var Customers = make([]*Customer, 0)

// type Customer struct {
// 	CustomerId   uuid.UUID
// 	FirstName    string
// 	LastName     string
// 	UserName     string
// 	Password     string
// 	TotalBalance int
// 	IsAdmin      bool
// 	IsActive     bool
// 	Accounts     []*account_service.Account
// }

// func NewCustomer(FirstName, LastName string, UserName, Password string, IsAdmin bool) *Customer {
// 	hashedPassword, _ := utils.GenerateHash(Password)
// 	var newCustomerObject *Customer = &Customer{
// 		CustomerId:   uuid.New(),
// 		FirstName:    FirstName,
// 		LastName:     LastName,
// 		UserName:     UserName,
// 		Password:     string(hashedPassword),
// 		TotalBalance: 0,
// 		IsAdmin:      IsAdmin,
// 		IsActive:     true,
// 		Accounts:     make([]*account_service.Account, 0),
// 	}
// 	Customers = append(Customers, newCustomerObject)
// 	return newCustomerObject
// }

// // CRUD OPERATIONS ON USERS

// func (c *Customer) CreateCustomer(FirstName, LastName string, UserName, Password string) (customer *Customer) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	customer = NewCustomer(FirstName, LastName, UserName, Password, false)
// 	panic(guru_errors.NewUserError(guru_errors.CreatedUser).GetSpecificMessage())

// }
// func CreateAdmin(FirstName, LastName string, UserName, Password string) (customer *Customer) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	customer = NewCustomer(FirstName, LastName, UserName, Password, true)
// 	panic(guru_errors.NewAdminError(guru_errors.CreatedAdmin).GetSpecificMessage())
// }

// // func (c *Customer) AddAccount(bankId uuid.UUID, balance int) (flag bool, accountObject *account_service.Account) {
// // 	flag = false
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()
// // 	if c.IsActive {
// // 		accountObject = account_service.CreateAccount(bankId, c.CustomerId, balance)

// // 		c.Accounts = append(c.Accounts, accountObject)
// // 		flag = true
// // 		return flag, accountObject
// // 	}
// // 	panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())

// // }

// func (c *Customer) ReadAllCustomers() (allCustomers []*Customer) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(Customers); i++ {
// 		if Customers[i].IsActive {
// 			allCustomers = append(allCustomers, Customers[i])
// 		}
// 	}

// 	panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())

// }

// func (c *Customer) ReadCustomerById(customerIdTemp uuid.UUID) (requiredCustomer *Customer) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)

// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	var requiredCustomerTemp *Customer
// 	for i := 0; i < len(Customers); i++ {
// 		if Customers[i].CustomerId == customerIdTemp {
// 			requiredCustomerTemp = Customers[i]
// 			break
// 		}
// 	}
// 	if requiredCustomerTemp.IsActive {
// 		requiredCustomer = requiredCustomerTemp
// 		panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())

// }
// func (c *Customer) UpdateCustomer(customerIdTemp uuid.UUID, updateField, updateValue interface{}) (updatedCustomer *Customer) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)

// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	var requiredCustomer *Customer = c.ReadCustomerById(customerIdTemp)
// 	if requiredCustomer == nil {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	switch updateValue := updateValue.(type) {
// 	case string:
// 		if updateField == "FirstName" {
// 			requiredCustomer.FirstName = updateValue
// 		} else {
// 			requiredCustomer.LastName = updateValue
// 		}

// 	}
// 	updatedCustomer = requiredCustomer

// 	panic(guru_errors.NewUserError(guru_errors.UpdatedUser).GetSpecificMessage())

// }
// func (c *Customer) UpdateCustomerObject(customerIdTemp uuid.UUID, customerTempObject *Customer) (updatedCustomer *Customer) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)

// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	var requiredCustomer *Customer = c.ReadCustomerById(customerIdTemp)
// 	if requiredCustomer == nil {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	if customerTempObject.FirstName != "" && requiredCustomer.FirstName != customerTempObject.FirstName {
// 		requiredCustomer.FirstName = customerTempObject.FirstName
// 	}
// 	if customerTempObject.LastName != "" && requiredCustomer.LastName != customerTempObject.LastName {
// 		requiredCustomer.LastName = customerTempObject.LastName
// 	}
// 	if customerTempObject.UserName != "" && requiredCustomer.UserName != customerTempObject.UserName {
// 		requiredCustomer.UserName = customerTempObject.UserName
// 	}

// 	updatedCustomer = requiredCustomer

// 	panic(guru_errors.NewUserError(guru_errors.UpdatedUser).GetSpecificMessage())

// }
// func (c *Customer) DeleteCustomer(customerIdTemp uuid.UUID) (deletedCustomer *Customer) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	var requiredCustomer *Customer = c.ReadCustomerById(customerIdTemp)
// 	if requiredCustomer == nil {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	for i := 0; i < len(requiredCustomer.Accounts); i++ {
// 		requiredCustomer.Accounts[i].DeleteAccount()
// 	}
// 	requiredCustomer.SetIsActive()
// 	deletedCustomer = requiredCustomer
// 	panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())

// }

// // CRUD OPERATIONS ON BANKS

// func (c *Customer) CreateBank(fullName string) (bank *bank_service.Bank) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	bank = bank_service.CreateBank(fullName)

// 	panic(guru_errors.NewBankError(guru_errors.CreatedBank).GetSpecificMessage())

// }

// func (c *Customer) ReadBankById(bankIdTemp uuid.UUID) (bank *bank_service.Bank) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	flag, requiredBank := bank_service.ReadBankById(bankIdTemp)
// 	if !flag {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())

// 	}
// 	bank = requiredBank
// 	panic(guru_errors.NewBankError(guru_errors.ReadBank).GetSpecificMessage())

// }

// func (c *Customer) ReadAllBanks() (allBanksInfo []*bank_service.Bank) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	allBanksInfo = bank_service.ReadAllBanks()
// 	panic(guru_errors.NewBankError(guru_errors.ReadBank).GetSpecificMessage())

// }
// func (c *Customer) UpdateBank(bankIdTemp uuid.UUID, updateValue string) (updatedBank *bank_service.Bank) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	var requiredBank *bank_service.Bank = c.ReadBankById(bankIdTemp)
// 	if requiredBank == nil {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}
// 	updatedBank = requiredBank.UpdateBank(updateValue)

// 	panic(guru_errors.NewBankError(guru_errors.UpdatedBank).GetSpecificMessage())

// }

// func (c *Customer) UpdateBankObject(bankIdTemp uuid.UUID, bankTempObject *bank_service.Bank) (updatedBank *bank_service.Bank) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	var requiredBank *bank_service.Bank = c.ReadBankById(bankIdTemp)
// 	if requiredBank == nil {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}
// 	updatedBank = requiredBank.UpdateBankObject(bankTempObject)

// 	panic(guru_errors.NewBankError(guru_errors.UpdatedBank).GetSpecificMessage())

// }
// func (c *Customer) DeleteBank(bankIdTemp uuid.UUID) (deletedBank *bank_service.Bank) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}

// 	var requiredBank *bank_service.Bank = c.ReadBankById(bankIdTemp)
// 	if requiredBank == nil {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}

// 	if requiredBank.CheckBankContainsActiveAccounts() {
// 		panic(guru_errors.NewBankError(guru_errors.BankContainsAccounts).GetSpecificMessage())
// 	}
// 	deletedBank = requiredBank.DeleteBank()
// 	panic(guru_errors.NewBankError(guru_errors.DeletedBank).GetSpecificMessage())

// }

// // CRUD OPERATIONS ON ACCOUNTS

// func (c *Customer) CreateAccount(bankIdTemp uuid.UUID, balance int) (account *account_service.Account) {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	if balance < 1000 {
// 		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
// 	}

// 	account = account_service.CreateAccount(bankIdTemp, c.CustomerId, balance)
// 	flag, requiredBank := bank_service.ReadBankById(bankIdTemp)
// 	if !flag {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}
// 	requiredBank.Accounts = append(requiredBank.Accounts, account)
// 	c.Accounts = append(c.Accounts, account)

// 	panic(guru_errors.NewAccountError(guru_errors.CreatedAccount).GetSpecificMessage())

// }
// func (c *Customer) ReadAccountById(accountIdTemp uuid.UUID) (requiredAccount *account_service.Account) {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	var requiredAccountTemp *account_service.Account
// 	for i := 0; i < len(c.Accounts); i++ {
// 		if c.Accounts[i].GetAccountNumber() == accountIdTemp {
// 			requiredAccountTemp = c.Accounts[i]
// 			break
// 		}
// 	}
// 	flag, requiredAccountPrint := requiredAccountTemp.ReadAccount()
// 	if !flag {

// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}
// 	requiredAccount = requiredAccountPrint
// 	panic(guru_errors.NewAccountError(guru_errors.ReadAccount).GetSpecificMessage())

// }
// func (c *Customer) ReadAllAccountsOfCustomer() (allAccounts []*account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	for i := 0; i < len(c.Accounts); i++ {
// 		flag, singleAccount := c.Accounts[i].ReadAccount()
// 		if flag {
// 			allAccounts = append(allAccounts, singleAccount)
// 		}
// 	}
// 	panic(guru_errors.NewAccountError(guru_errors.ReadAccount).GetSpecificMessage())

// }
// func (c *Customer) UpdateAccount(accountIdTemp uuid.UUID, updateField string, updateValue interface{}) (updatedAccount *account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	requiredAccount := c.ReadAccountById(accountIdTemp)
// 	if requiredAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}
// 	updatedAccountTemp := requiredAccount.UpdateAccount(updateField, updateValue)

// 	flag1, requiredBank := bank_service.ReadBankById(updatedAccountTemp.GetBankId())
// 	if !flag1 {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}

// 	requiredCustomer := ReadCustomerById(updatedAccountTemp.GetCustomerId())
// 	if requiredCustomer == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(requiredBank.Accounts); i++ {
// 		if requiredBank.Accounts[i].GetAccountNumber() == updatedAccountTemp.GetAccountNumber() {
// 			requiredBank.Accounts[i] = updatedAccountTemp
// 			break
// 		}
// 	}

// 	for i := 0; i < len(c.Accounts); i++ {
// 		if requiredCustomer.Accounts[i].GetAccountNumber() == updatedAccountTemp.GetAccountNumber() {
// 			requiredCustomer.Accounts[i] = updatedAccountTemp
// 			break
// 		}
// 	}
// 	updatedAccount = updatedAccountTemp

// 	panic(guru_errors.NewAccountError(guru_errors.UpdatedAccount).GetSpecificMessage())
// }
// func (c *Customer) DeleteAccount(accountIdTemp uuid.UUID) (deletedAccount *account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	requiredAccount := c.ReadAccountById(accountIdTemp)
// 	if requiredAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}
// 	deletedAccountTemp := requiredAccount.DeleteAccount()

// 	flag1, requiredBank := bank_service.ReadBankById(deletedAccountTemp.GetBankId())
// 	if !flag1 {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}
// 	requiredCustomer := ReadCustomerById(deletedAccountTemp.GetCustomerId())
// 	if requiredCustomer == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(requiredBank.Accounts); i++ {
// 		if requiredBank.Accounts[i].GetAccountNumber() == deletedAccountTemp.GetAccountNumber() {
// 			requiredBank.Accounts[i] = deletedAccountTemp
// 			break
// 		}
// 	}

// 	for i := 0; i < len(requiredCustomer.Accounts); i++ {
// 		if requiredCustomer.Accounts[i].GetAccountNumber() == deletedAccountTemp.GetAccountNumber() {
// 			requiredCustomer.Accounts[i] = deletedAccountTemp
// 			break
// 		}
// 	}

// 	deletedAccount = deletedAccountTemp
// 	panic(guru_errors.NewAccountError(guru_errors.DeletedAccount).GetSpecificMessage())

// }

// // HELPER FUNCTIONS

// func (c *Customer) DepositMoney(accountNumberTemp uuid.UUID, amount int) (updatedAccount *account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	var requiredAccount *account_service.Account = c.ReadAccountById(accountNumberTemp)
// 	if requiredAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}
// 	depositedAccount := requiredAccount.DepositMoney(amount)
// 	updatedAccountTemp := c.UpdateAccount(depositedAccount.GetAccountNumber(), "balance", depositedAccount.GetBalance())

// 	if updatedAccountTemp == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.MoneyDepositedError).GetSpecificMessage())
// 	}
// 	updatedAccount = updatedAccountTemp
// 	panic(guru_errors.NewAccountError(guru_errors.MoneyDeposited).GetSpecificMessage())

// }

// func (c *Customer) WithdrawMoney(accountNumberTemp uuid.UUID, amount int) (updatedAccount *account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	var requiredAccount *account_service.Account = c.ReadAccountById(accountNumberTemp)
// 	if requiredAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}

// 	if requiredAccount.GetBalance()-amount < 0 {
// 		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
// 	}
// 	withdrawAccount := requiredAccount.WithdrawMoney(amount)

// 	updatedAccountTemp := c.UpdateAccount(withdrawAccount.GetAccountNumber(), "balance", withdrawAccount.GetBalance())
// 	if updatedAccountTemp == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.MoneyWithdrawError).GetSpecificMessage())
// 	}
// 	updatedAccount = updatedAccountTemp
// 	panic(guru_errors.NewAccountError(guru_errors.MoneyWithdraw).GetSpecificMessage())
// }

// func (c *Customer) TransferMoney(senderAccountNumber, receiverAccountNumber uuid.UUID, amount int) (updatedSenderAccount *account_service.Account, updatedReceiverAccount *account_service.Account) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	var requiredSenderAccount *account_service.Account = c.ReadAccountById(senderAccountNumber)
// 	var requiredReceiverAccount *account_service.Account

// 	for i := 0; i < len(bank_service.Banks); i++ {
// 		for j := 0; j < len(bank_service.Banks[i].Accounts); j++ {
// 			if bank_service.Banks[i].Accounts[j].GetAccountNumber() == receiverAccountNumber {
// 				requiredReceiverAccount = bank_service.Banks[i].Accounts[j]
// 				break
// 			}
// 		}

// 	}

// 	if requiredSenderAccount == nil || requiredReceiverAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}
// 	if requiredSenderAccount.GetBalance()-amount < 0 {
// 		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
// 	}
// 	transferSenderAccount, transferReceiverAccount := requiredSenderAccount.TransferMoney(requiredReceiverAccount, amount)

// 	requiredReceiverCustomer := ReadCustomerById(requiredReceiverAccount.GetCustomerId())
// 	if requiredReceiverAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	updatedSenderAccountTemp := c.UpdateAccount(transferSenderAccount.GetAccountNumber(), "balance", transferSenderAccount.GetBalance())
// 	updatedReceiverAccountTemp := requiredReceiverCustomer.UpdateAccount(transferReceiverAccount.GetAccountNumber(), "balance", transferReceiverAccount.GetBalance())

// 	if updatedSenderAccountTemp == nil || updatedReceiverAccountTemp == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.MoneyTransferedError).GetSpecificMessage())
// 	}
// 	//
// 	flag1, senderBank := bank_service.ReadBankById(updatedSenderAccountTemp.GetBankId())
// 	flag2, receiverBank := bank_service.ReadBankById(updatedReceiverAccountTemp.GetBankId())
// 	if !flag1 {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}
// 	if !flag2 {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}

// 	senderBank.GetBankPassbook().AddEntry(senderBank.GetBankId(), receiverBank.GetBankId(), 0-amount)
// 	receiverBank.GetBankPassbook().AddEntry(senderBank.GetBankId(), receiverBank.GetBankId(), amount)

// 	updatedSenderAccount = updatedSenderAccountTemp
// 	updatedReceiverAccount = updatedReceiverAccountTemp
// 	panic(guru_errors.NewAccountError(guru_errors.MoneyTransfered).GetSpecificMessage())

// }

// func ReadCustomerById(customerIdTemp uuid.UUID) (requiredCustomer *Customer) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	var customer *Customer
// 	for i := 0; i < len(Customers); i++ {
// 		if Customers[i].GetCustomerId() == customerIdTemp {
// 			customer = Customers[i]
// 			break
// 		}
// 	}

// 	if !customer.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	requiredCustomer = customer
// 	panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())
// }
// func (c *Customer) GetTotalBalance() (TotalBalance int) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(c.Accounts); i++ {
// 		TotalBalance += c.Accounts[i].GetBalance()
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserTotalBalance).GetSpecificMessage())
// }

// func (c *Customer) GetAllIndividualAccountBalance() (mapAccountBalance map[uuid.UUID]int) {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}

// 	mapAccountBalance = make(map[uuid.UUID]int, 0)

// 	for i := 0; i < len(c.Accounts); i++ {
// 		mapAccountBalance[c.Accounts[i].GetAccountNumber()] = c.Accounts[i].GetBalance()
// 	}

// 	panic(guru_errors.NewUserError(guru_errors.UserAccoutBalanceMap).GetSpecificMessage())

// }

// func (c *Customer) GetNetWorthOfEachBank() (mapBankBalance map[uuid.UUID]int) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	mapBankBalance = make(map[uuid.UUID]int, 0)
// 	for i := 0; i < len(bank_service.Banks); i++ {
// 		mapBankBalance[bank_service.Banks[i].GetBankId()] = bank_service.Banks[i].GetNetWorthOfBank()
// 	}

// 	panic(guru_errors.NewBankError(guru_errors.BankNetWorthMap).GetSpecificMessage())
// }

// func (c *Customer) GetNetWorthOfGivenBank(bankIdTemp uuid.UUID) (mapBankBalance map[uuid.UUID]int) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	mapBankBalance = make(map[uuid.UUID]int, 0)
// 	for i := 0; i < len(bank_service.Banks); i++ {
// 		if bank_service.Banks[i].GetBankId() == bankIdTemp {
// 			mapBankBalance[bank_service.Banks[i].GetBankId()] = bank_service.Banks[i].GetNetWorthOfBank()
// 		}
// 	}

// 	panic(guru_errors.NewBankError(guru_errors.BankNetWorthMap).GetSpecificMessage())
// }

// func (c *Customer) GetPassbookInRange(accountNumberTemp uuid.UUID, startDate string, endDate string) (requiredPassbookInRange *guru_passbook.Passbook) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	if !c.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	requiredAccount := c.ReadAccountById(accountNumberTemp)
// 	if requiredAccount == nil {
// 		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
// 	}

// 	var startDateGoTime time.Time
// 	var endDateGoTime time.Time

// 	if startDate == "" {
// 		startDate = "1970-01-01"
// 	}
// 	if endDate == "" {
// 		endDate = time.Now().Format("2006-01-02")
// 	}
// 	startDateGoTime, _ = time.Parse("2006-01-02", startDate)
// 	endDateGoTime, _ = time.Parse("2006-01-02", endDate)

// 	// fmt.Println("GetPassbookInRange Func : ", startDateGoTime, "    ", endDateGoTime)
// 	requiredPassbookInRange = requiredAccount.GetPassbook(startDateGoTime, endDateGoTime)
// 	panic(guru_errors.NewAccountError(guru_errors.PassbookReadInRange).GetSpecificMessage())

// }

// func (c *Customer) BankTransferMapNameBalanceByBankId(bankIdTemp uuid.UUID, fromDate string, toDate string) (bankTransferAllMapByBankId map[string]int) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	var requiredBank *bank_service.Bank = c.ReadBankById(bankIdTemp)
// 	if requiredBank == nil {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}

// 	var fromDateGoTime time.Time
// 	var toDateGoTime time.Time

// 	if fromDate == "" {
// 		fromDate = "1970-01-01"
// 	}
// 	if toDate == "" {
// 		toDate = time.Now().Format("2006-01-02")
// 	}
// 	fromDateGoTime, _ = time.Parse("2006-01-02", fromDate)
// 	toDateGoTime, _ = time.Parse("2006-01-02", toDate)

// 	bankTransferTemp := requiredBank.ReadPassbookFromRange(fromDateGoTime, toDateGoTime)
// 	bankTransferAllMapByBankId = make(map[string]int, 0)
// 	for key, value := range bankTransferTemp {
// 		bankTransferAllMapByBankId[GetBankNameById(key)] = value
// 	}
// 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
// }
// func (c *Customer) BankTransferMapNameBalanceAll(fromDate string, toDate string) (bankTransferAllMap map[string]map[string]int) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.IsAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.IsActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	bankTransferAllMap = make(map[string]map[string]int)
// 	for i := 0; i < len(bank_service.Banks); i++ {
// 		bankTransferAllMapByBankId := c.BankTransferMapNameBalanceByBankId(bank_service.Banks[i].GetBankId(), fromDate, toDate)
// 		bankTransferAllMap[bank_service.Banks[i].GetBankName()] = bankTransferAllMapByBankId
// 	}

// 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
// }

// // func (c *Customer) BankTransferAllMapByBankId(bankIdTemp uuid.UUID) (bankTransferAllMapByBankId map[string]int) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()

// // 	if !c.IsAdmin {
// // 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// // 	}
// // 	if !c.IsActive {
// // 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// // 	}
// // 	var requiredBank *bank_service.Bank = c.ReadBankById(bankIdTemp)
// // 	if requiredBank == nil {
// // 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// // 	}

// // 	var bankTransferAllMapIdBalance map[uuid.UUID]int = requiredBank.GetBankTransferAllMap()
// // 	for key, value := range bankTransferAllMapIdBalance {
// // 		bankTransferAllMapByBankId[c.GetBankNameById(key)] = value
// // 	}
// // 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())

// // }

// // func (c *Customer) BankTransferAllMapAll() (bankTransferAllMap []map[string]int) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()

// // 	if !c.IsAdmin {
// // 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// // 	}
// // 	if !c.IsActive {
// // 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// // 	}
// // 	for i := 0; i < len(bank_service.Banks); i++ {
// // 		bankTransferAllMap = append(bankTransferAllMap, c.BankTransferAllMapByBankId(bank_service.Banks[i].GetBankId()))

// // 	}
// // 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
// // }

// func (c *Customer) GetBankNameById(bankIdTemp uuid.UUID) string {
// 	return c.ReadBankById(bankIdTemp).GetBankName()

// }
// func GetBankNameById(bankIdTemp uuid.UUID) (bankName string) {
// 	for i := 0; i < len(bank_service.Banks); i++ {
// 		if bank_service.Banks[i].GetBankId() == bankIdTemp {
// 			bankName = bank_service.Banks[i].GetBankName()
// 			break
// 		}

// 	}
// 	return bankName
// }

// func ReadCustomerByUserName(customerUserNameTemp string) (requiredCustomer *Customer) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	var requiredCustomerTemp *Customer
// 	for i := 0; i < len(Customers); i++ {
// 		if Customers[i].UserName == customerUserNameTemp {
// 			requiredCustomerTemp = Customers[i]
// 			break
// 		}
// 	}
// 	if !requiredCustomerTemp.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
// 	}
// 	requiredCustomer = requiredCustomerTemp

// 	panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())

// }
// func (c *Customer) GetCustomerId() uuid.UUID {
// 	return c.CustomerId
// }
// func (c *Customer) SetIsActive() {
// 	c.IsActive = false
// }
// func (c *Customer) GetIsActive() bool {
// 	return c.IsActive
// }
// func (c *Customer) GetIsAdmin() bool {
// 	return c.IsAdmin
// }
