package guru_customer

import (
	"bankingapp/guru_account"
	"bankingapp/guru_bank"
	"bankingapp/guru_errors"
	"bankingapp/guru_passbook"
	"fmt"
	"time"

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
	if !flag {
		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())

	}
	bank = requiredBank
	panic(guru_errors.NewBankError(guru_errors.ReadBank).GetSpecificMessage())

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
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
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
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
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
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
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
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	requiredAccount := c.ReadAccountById(accountIdTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}
	updatedAccountTemp := requiredAccount.UpdateAccount(updateField, updateValue)

	flag1, requiredBank := guru_bank.ReadBankById(updatedAccountTemp.GetBankId())
	if !flag1 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}
	flag2, requiredCustomer := ReadCustomerById(updatedAccountTemp.GetCustomerId())
	if !flag2 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	for i := 0; i < len(requiredBank.Accounts); i++ {
		if requiredBank.Accounts[i].GetAccountNumber() == updatedAccountTemp.GetAccountNumber() {
			requiredBank.Accounts[i] = updatedAccountTemp
			break
		}
	}

	for i := 0; i < len(c.Accounts); i++ {
		if requiredCustomer.Accounts[i].GetAccountNumber() == updatedAccountTemp.GetAccountNumber() {
			requiredCustomer.Accounts[i] = updatedAccountTemp
			break
		}
	}
	updatedAccount = updatedAccountTemp

	panic(guru_errors.NewAccountError(guru_errors.UpdatedAccount).GetSpecificMessage())
}
func (c *Customer) DeleteAccount(accountIdTemp uuid.UUID) (deletedAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	requiredAccount := c.ReadAccountById(accountIdTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}
	deletedAccountTemp := requiredAccount.DeleteAccount()

	flag1, requiredBank := guru_bank.ReadBankById(deletedAccountTemp.GetBankId())
	if !flag1 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}
	flag2, requiredCustomer := ReadCustomerById(deletedAccountTemp.GetCustomerId())
	if !flag2 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	for i := 0; i < len(requiredBank.Accounts); i++ {
		if requiredBank.Accounts[i].GetAccountNumber() == deletedAccountTemp.GetAccountNumber() {
			requiredBank.Accounts[i] = deletedAccountTemp
			break
		}
	}

	for i := 0; i < len(requiredCustomer.Accounts); i++ {
		if requiredCustomer.Accounts[i].GetAccountNumber() == deletedAccountTemp.GetAccountNumber() {
			requiredCustomer.Accounts[i] = deletedAccountTemp
			break
		}
	}

	deletedAccount = deletedAccountTemp
	panic(guru_errors.NewAccountError(guru_errors.DeletedAccount).GetSpecificMessage())

}

// HELPER FUNCTIONS

func (c *Customer) DepositMoney(accountNumberTemp uuid.UUID, amount int) (updatedAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	var requiredAccount *guru_account.Account = c.ReadAccountById(accountNumberTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}
	depositedAccount := requiredAccount.DepositMoney(amount)
	updatedAccountTemp := c.UpdateAccount(depositedAccount.GetAccountNumber(), "balance", depositedAccount.GetBalance())

	if updatedAccountTemp == nil {
		panic(guru_errors.NewAccountError(guru_errors.MoneyDepositedError).GetSpecificMessage())
	}
	updatedAccount = updatedAccountTemp
	panic(guru_errors.NewAccountError(guru_errors.MoneyDeposited).GetSpecificMessage())

}

func (c *Customer) WithdrawMoney(accountNumberTemp uuid.UUID, amount int) (updatedAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	var requiredAccount *guru_account.Account = c.ReadAccountById(accountNumberTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}

	if requiredAccount.GetBalance()-amount < 0 {
		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
	}
	withdrawAccount := requiredAccount.WithdrawMoney(amount)

	updatedAccountTemp := c.UpdateAccount(withdrawAccount.GetAccountNumber(), "balance", withdrawAccount.GetBalance())
	if updatedAccountTemp == nil {
		panic(guru_errors.NewAccountError(guru_errors.MoneyWithdrawError).GetSpecificMessage())
	}
	updatedAccount = updatedAccountTemp
	panic(guru_errors.NewAccountError(guru_errors.MoneyWithdraw).GetSpecificMessage())
}

func (c *Customer) TransferMoney(senderAccountNumber, receiverAccountNumber uuid.UUID, amount int) (updatedSenderAccount *guru_account.Account, updatedReceiverAccount *guru_account.Account) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	var requiredSenderAccount *guru_account.Account = c.ReadAccountById(senderAccountNumber)
	var requiredReceiverAccount *guru_account.Account

	for i := 0; i < len(guru_bank.Banks); i++ {
		for j := 0; j < len(guru_bank.Banks[i].Accounts); j++ {
			if guru_bank.Banks[i].Accounts[j].GetAccountNumber() == receiverAccountNumber {
				requiredReceiverAccount = guru_bank.Banks[i].Accounts[j]
				break
			}
		}

	}

	if requiredSenderAccount == nil || requiredReceiverAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}
	if requiredSenderAccount.GetBalance()-amount < 0 {
		panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
	}
	transferSenderAccount, transferReceiverAccount := requiredSenderAccount.TransferMoney(requiredReceiverAccount, amount)

	flag, requiredReceiverCustomer := ReadCustomerById(requiredReceiverAccount.GetCustomerId())
	if !flag {
		panic(guru_errors.NewAccountError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	updatedSenderAccountTemp := c.UpdateAccount(transferSenderAccount.GetAccountNumber(), "balance", transferSenderAccount.GetBalance())
	updatedReceiverAccountTemp := requiredReceiverCustomer.UpdateAccount(transferReceiverAccount.GetAccountNumber(), "balance", transferReceiverAccount.GetBalance())

	if updatedSenderAccountTemp == nil || updatedReceiverAccountTemp == nil {
		panic(guru_errors.NewAccountError(guru_errors.MoneyTransferedError).GetSpecificMessage())
	}
	//
	flag1, senderBank := guru_bank.ReadBankById(updatedSenderAccountTemp.GetBankId())
	flag2, receiverBank := guru_bank.ReadBankById(updatedReceiverAccountTemp.GetBankId())
	if !flag1 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}
	if !flag2 {
		panic(guru_errors.NewAccountError(guru_errors.DeletedBankStatus).GetSpecificMessage())
	}

	senderBank.GetBankPassbook().AddEntry(senderBank.GetBankId(), receiverBank.GetBankId(), 0-amount)
	receiverBank.GetBankPassbook().AddEntry(senderBank.GetBankId(), receiverBank.GetBankId(), amount)

	updatedSenderAccount = updatedSenderAccountTemp
	updatedReceiverAccount = updatedReceiverAccountTemp
	panic(guru_errors.NewAccountError(guru_errors.MoneyTransfered).GetSpecificMessage())

}

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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	for i := 0; i < len(c.Accounts); i++ {
		totalBalance += c.Accounts[i].GetBalance()
	}
	panic(guru_errors.NewUserError(guru_errors.UserTotalBalance).GetSpecificMessage())
}

func (c *Customer) GetAllIndividualAccountBalance() (mapAccountBalance map[uuid.UUID]int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}

	mapAccountBalance = make(map[uuid.UUID]int, 0)

	for i := 0; i < len(c.Accounts); i++ {
		mapAccountBalance[c.Accounts[i].GetAccountNumber()] = c.Accounts[i].GetBalance()
	}

	panic(guru_errors.NewUserError(guru_errors.UserAccoutBalanceMap).GetSpecificMessage())

}

func (c *Customer) GetNetWorthOfEachBank() (mapBankBalance map[uuid.UUID]int) {
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
	for i := 0; i < len(guru_bank.Banks); i++ {
		mapBankBalance[guru_bank.Banks[i].GetBankId()] = guru_bank.Banks[i].GetNetWorthOfBank()
	}

	panic(guru_errors.NewBankError(guru_errors.BankNetWorthMap).GetSpecificMessage())
}

func (c *Customer) GetPassbookInRange(accountNumberTemp uuid.UUID, startDate string, endDate string) (requiredPassbookInRange *guru_passbook.Passbook) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if !c.isActive {
		panic(guru_errors.NewUserError(guru_errors.DeletedUserStatus).GetSpecificMessage())
	}
	requiredAccount := c.ReadAccountById(accountNumberTemp)
	if requiredAccount == nil {
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())
	}

	var startDateGoTime time.Time
	var endDateGoTime time.Time

	if startDate == "" {
		startDate = "1970-01-01"
	}
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}
	startDateGoTime, _ = time.Parse("2006-01-02", startDate)
	endDateGoTime, _ = time.Parse("2006-01-02", endDate)
	requiredPassbookInRange = requiredAccount.GetPassbook(startDateGoTime, endDateGoTime)
	panic(guru_errors.NewAccountError(guru_errors.PassbookReadInRange).GetSpecificMessage())

}

func (c *Customer) BankTransferMapNameBalanceByBankId(bankIdTemp uuid.UUID, fromDate string, toDate string) (bankTransferAllMapByBankId map[string]int) {
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

	var fromDateGoTime time.Time
	var toDateGoTime time.Time

	if fromDate == "" {
		fromDate = "1970-01-01"
	}
	if toDate == "" {
		toDate = time.Now().Format("2006-01-02")
	}
	fromDateGoTime, _ = time.Parse("2006-01-02", fromDate)
	toDateGoTime, _ = time.Parse("2006-01-02", toDate)

	bankTransferTemp := requiredBank.ReadPassbookFromRange(fromDateGoTime, toDateGoTime)
	for key, value := range bankTransferTemp {
		bankTransferAllMapByBankId[c.GetBankNameById(key)] = value
	}
	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
}
func (c *Customer) BankTransferMapNameBalanceAll(fromDate string, toDate string) (bankTransferAllMap map[string]map[string]int) {
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

	for i := 0; i < len(guru_bank.Banks); i++ {
		bankTransferAllMapByBankId := c.BankTransferMapNameBalanceByBankId(guru_bank.Banks[i].GetBankId(), fromDate, toDate)
		bankTransferAllMap[guru_bank.Banks[i].GetBankName()] = bankTransferAllMapByBankId
	}

	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
}

// func (c *Customer) BankTransferAllMapByBankId(bankIdTemp uuid.UUID) (bankTransferAllMapByBankId map[string]int) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.isAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.isActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	var requiredBank *guru_bank.Bank = c.ReadBankById(bankIdTemp)
// 	if requiredBank == nil {
// 		panic(guru_errors.NewBankError(guru_errors.DeletedBankStatus).GetSpecificMessage())
// 	}

// 	var bankTransferAllMapIdBalance map[uuid.UUID]int = requiredBank.GetBankTransferAllMap()
// 	for key, value := range bankTransferAllMapIdBalance {
// 		bankTransferAllMapByBankId[c.GetBankNameById(key)] = value
// 	}
// 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())

// }

// func (c *Customer) BankTransferAllMapAll() (bankTransferAllMap []map[string]int) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !c.isAdmin {
// 		panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
// 	}
// 	if !c.isActive {
// 		panic(guru_errors.NewAdminError(guru_errors.DeletedAdmin).GetSpecificMessage())
// 	}
// 	for i := 0; i < len(guru_bank.Banks); i++ {
// 		bankTransferAllMap = append(bankTransferAllMap, c.BankTransferAllMapByBankId(guru_bank.Banks[i].GetBankId()))

// 	}
// 	panic(guru_errors.NewBankError(guru_errors.ReadBankTransferAllMap).GetSpecificMessage())
// }

func (c *Customer) GetBankNameById(bankIdTemp uuid.UUID) string {
	return c.ReadBankById(bankIdTemp).GetBankName()

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
