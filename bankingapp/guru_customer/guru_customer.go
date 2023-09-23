package guru_customer

import (
	"bankingapp/guru_account"
	"bankingapp/guru_bank"
	"bankingapp/guru_errors"
	"fmt"

	"github.com/google/uuid"
)

var Customers = make([]*Customer, 0)

type Customer struct {
	customerId   uuid.UUID
	firstName    string
	lastName     string
	totalBalance int
	isAdmin      bool
	isActive     bool
	Accounts     []*guru_account.Account
}

func NewCustomer(firstName, lastName string, isAdmin bool) *Customer {
	var newCustomerObject *Customer = &Customer{
		customerId:   uuid.New(),
		firstName:    firstName,
		lastName:     lastName,
		totalBalance: 0,
		isAdmin:      isAdmin,
		isActive:     true,
		Accounts:     make([]*guru_account.Account, 0),
	}
	Customers = append(Customers, newCustomerObject)
	return newCustomerObject
}

// CRUD OPERATIONS ON USERS

func (c *Customer) CreateCustomer(firstName, lastName string) (customer *Customer) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	customer = NewCustomer(firstName, lastName, false)
	panic(guru_errors.NewUserError(guru_errors.CreatedUser).GetSpecificMessage())

}
func CreateAdmin(firstName, lastName string) (customer *Customer) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	customer = NewCustomer(firstName, lastName, true)
	panic(guru_errors.NewAdminError(guru_errors.CreatedAdmin).GetSpecificMessage())
}

// func (c *Customer) AddAccount(bankId uuid.UUID, balance int) (flag bool, accountObject *guru_account.Account) {
// 	flag = false
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if c.isActive {
// 		accountObject = guru_account.CreateAccount(bankId, c.customerId, balance)

// 		c.Accounts = append(c.Accounts, accountObject)
// 		flag = true
// 		return flag, accountObject
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())

// }

func (c *Customer) ReadAllCustomers() (allCustomers []*Customer) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	for i := 0; i < len(Customers); i++ {
		if Customers[i].isActive {
			allCustomers = append(allCustomers, Customers[i])
		}
	}

	panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())

}

func (c *Customer) ReadCustomerById(customerIdTemp uuid.UUID) (requiredCustomer *Customer) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	var requiredCustomerTemp *Customer
	for i := 0; i < len(Customers); i++ {
		if Customers[i].customerId == customerIdTemp {
			requiredCustomerTemp = Customers[i]
			break
		}
	}
	if requiredCustomerTemp.isActive {
		requiredCustomer = requiredCustomerTemp
		panic(guru_errors.NewUserError(guru_errors.ReadUser).GetSpecificMessage())
	}
	panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())

}
func (c *Customer) UpdateCustomer(customerIdTemp uuid.UUID, updateField, updateValue interface{}) (updatedCustomer *Customer) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}
	var requiredCustomer *Customer = c.ReadCustomerById(customerIdTemp)
	if requiredCustomer == nil {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	switch updateValue := updateValue.(type) {
	case string:
		if updateField == "firstName" {
			requiredCustomer.firstName = updateValue
		} else {
			requiredCustomer.lastName = updateValue
		}

	}
	updatedCustomer = requiredCustomer

	panic(guru_errors.NewUserError(guru_errors.UpdatedUser).GetSpecificMessage())

}
func (c *Customer) DeleteCustomer(customerIdTemp uuid.UUID) (deletedCustomer *Customer) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	var requiredCustomer *Customer = c.ReadCustomerById(customerIdTemp)
	if requiredCustomer == nil {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	for i := 0; i < len(requiredCustomer.Accounts); i++ {
		requiredCustomer.Accounts[i].DeleteAccount()
	}
	requiredCustomer.SetIsActve()
	deletedCustomer = requiredCustomer
	panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())

}

// CRUD OPERATIONS ON BANKS

func (c *Customer) CreateBank(fullName string) (bank *guru_bank.Bank) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	bank = guru_bank.CreateBank(fullName)

	panic(guru_errors.NewBankError(guru_errors.CreatedBank).GetSpecificMessage())

}

func (c *Customer) ReadBankById(bankIdTemp uuid.UUID) (bank *guru_bank.Bank) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	flag, requiredBank := guru_bank.ReadBankById(bankIdTemp)
	if flag {
		bank = requiredBank
		panic(guru_errors.NewBankError(guru_errors.ReadBank).GetSpecificMessage())
	}

	panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())

}

func (c *Customer) ReadAllBanks() (allBanksInfo []*guru_bank.Bank) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	allBanksInfo = guru_bank.ReadAllBanks()
	panic(guru_errors.NewBankError(guru_errors.ReadBank).GetSpecificMessage())

}
func (c *Customer) UpdateBank(bankIdTemp uuid.UUID, updateValue string) (updatedBank *guru_bank.Bank) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	var requiredBank *guru_bank.Bank = c.ReadBankById(bankIdTemp)
	if requiredBank == nil {
		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}
	updatedBank = requiredBank.UpdateBank(updateValue)

	panic(guru_errors.NewBankError(guru_errors.UpdatedBank).GetSpecificMessage())

}
func (c *Customer) DeleteBank(bankIdTemp uuid.UUID) (deletedBank *guru_bank.Bank) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !c.isAdmin {
		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
	}
	if !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
	}

	var requiredBank *guru_bank.Bank = c.ReadBankById(bankIdTemp)
	if requiredBank == nil {
		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}

	if requiredBank.CheckBankContainsActiveAccounts() {
		panic(guru_errors.NewBankError(guru_errors.BankContainsAccounts).GetSpecificMessage())
	}
	deletedBank = requiredBank.DeleteBank()
	panic(guru_errors.NewBankError(guru_errors.DeletedBank).GetSpecificMessage())

}

// CRUD OPERATIONS ON ACCOUNTS

func (c *Customer) CreateAccount(bankIdTemp uuid.UUID, balance int) (account *guru_account.Account) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())
	}
	if balance < 1000 {
		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
	}

	account = guru_account.CreateAccount(bankIdTemp, c.customerId, balance)
	flag, requiredBank := guru_bank.ReadBankById(bankIdTemp)
	if !flag {
		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}
	requiredBank.Accounts = append(requiredBank.Accounts, account)
	c.Accounts = append(c.Accounts, account)

	panic(guru_errors.NewAccountError(guru_errors.CreatedAccount).GetSpecificMessage())

}
func (c *Customer) ReadAccountById(accountIdTemp uuid.UUID) (requiredAccount *guru_account.Account) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())
	}

	var requiredAccountTemp *guru_account.Account
	for i := 0; i < len(c.Accounts); i++ {
		if c.Accounts[i].GetAccountNumber() == accountIdTemp {
			requiredAccountTemp = c.Accounts[i]
			break
		}
	}
	flag, requiredAccountPrint := requiredAccountTemp.ReadAccount()
	if !flag {

		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}
	requiredAccount = requiredAccountPrint
	panic(guru_errors.NewAccountError(guru_errors.ReadAccount).GetSpecificMessage())

}
func (c *Customer) ReadAllAccountsOfCustomer() (allAccounts []*guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())
	}
	for i := 0; i < len(c.Accounts); i++ {
		flag, singleAccount := c.Accounts[i].ReadAccount()
		if flag {
			allAccounts = append(allAccounts, singleAccount)
		}
	}
	panic(guru_errors.NewAccountError(guru_errors.ReadAccount).GetSpecificMessage())

}
func (c *Customer) UpdateAccount(accountIdTemp uuid.UUID, updateField string, updateValue interface{}) (updatedAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())
	}

	requiredAccount := c.ReadAccountById(accountIdTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountAlready).GetSpecificMessage())
	}
	updatedAccount = requiredAccount.UpdateAccount(updateField, updateValue)

	_, requiredBank := guru_bank.ReadBankById(updatedAccount.GetBankId())

	for i := 0; i < len(requiredBank.Accounts); i++ {
		if requiredBank.Accounts[i].GetAccountNumber() == updatedAccount.GetAccountNumber() {
			requiredBank.Accounts[i] = updatedAccount
			break
		}
	}
	_, requiredCustomer := ReadCustomerById(updatedAccount.GetCustomerId())

	for i := 0; i < len(c.Accounts); i++ {
		if requiredCustomer.Accounts[i].GetAccountNumber() == updatedAccount.GetAccountNumber() {
			requiredCustomer.Accounts[i] = updatedAccount
			break
		}
	}

	panic(guru_errors.NewAccountError(guru_errors.UpdatedAccount).GetSpecificMessage())
}
func (c *Customer) DeleteAccount(accountIdTemp uuid.UUID) (deletedAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUser).GetSpecificMessage())
	}

	requiredAccount := c.ReadAccountById(accountIdTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountAlready).GetSpecificMessage())
	}
	deletedAccount = requiredAccount.DeleteAccount()

	_, requiredBank := guru_bank.ReadBankById(deletedAccount.GetBankId())

	for i := 0; i < len(requiredBank.Accounts); i++ {
		if requiredBank.Accounts[i].GetAccountNumber() == deletedAccount.GetAccountNumber() {
			requiredBank.Accounts[i] = deletedAccount
			break
		}
	}

	_, requiredCustomer := ReadCustomerById(deletedAccount.GetCustomerId())

	for i := 0; i < len(requiredCustomer.Accounts); i++ {
		if requiredCustomer.Accounts[i].GetAccountNumber() == deletedAccount.GetAccountNumber() {
			requiredCustomer.Accounts[i] = deletedAccount
			break
		}
	}

	panic(guru_errors.NewAccountError(guru_errors.DeletedAccount).GetSpecificMessage())

}

// HELPER FUNCTIONS

func ReadCustomerById(customerIdTemp uuid.UUID) (bool, *Customer) {
	var customer *Customer
	for i := 0; i < len(Customers); i++ {
		if Customers[i].GetCustomerId() == customerIdTemp {
			customer = Customers[i]
			break
		}
	}
	if customer.isActive {
		return true, customer
	}
	return false, customer
}
func (c *Customer) GetTotalBalance() (totalBalance int) {
	for i := 0; i < len(c.Accounts); i++ {
		totalBalance += c.Accounts[i].GetBalance()
	}
	return totalBalance
}

func (c *Customer) GetAllIndividualAccountBalance() map[uuid.UUID]int {

	var mapAccountBalance = make(map[uuid.UUID]int, 0)

	for i := 0; i < len(c.Accounts); i++ {
		mapAccountBalance[c.Accounts[i].GetAccountNumber()] = c.Accounts[i].GetBalance()
	}

	return mapAccountBalance

}

func (c *Customer) GetCustomerId() uuid.UUID {
	return c.customerId
}
func (c *Customer) SetIsActve() {
	c.isActive = false
}

func (c *Customer) GetIsActive() bool {
	return c.isActive
}
func (c *Customer) GetIsAdmin() bool {
	return c.isAdmin
}
