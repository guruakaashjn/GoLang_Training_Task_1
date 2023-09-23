package guru_account

import (
	"bankingapp/guru_errors"
	"bankingapp/guru_passbook"
	"fmt"

	"github.com/google/uuid"
)

type Account struct {
	accountNumber uuid.UUID
	bankId        uuid.UUID
	customerId    uuid.UUID
	isActive      bool
	balance       int
	passbook      *guru_passbook.Passbook
}

func NewAccount(bankId uuid.UUID, customerId uuid.UUID, balance int) *Account {
	accountNumber := uuid.New()
	var newPassbook *guru_passbook.Passbook = guru_passbook.NewPassbook(customerId, accountNumber, balance)
	return &Account{
		accountNumber: accountNumber,
		bankId:        bankId,
		customerId:    customerId,
		isActive:      true,
		balance:       balance,
		passbook:      newPassbook,
	}
}

func CreateAccount(bankId uuid.UUID, customerId uuid.UUID, balance int) (account *Account) {

	return NewAccount(bankId, customerId, balance)

}

func (a *Account) ReadAccount() (bool, *Account) {
	if a.isActive {
		return true, a
	}
	return false, a

}
func (a *Account) UpdateAccount(updateField string, updateValue interface{}) *Account {
	// balance, abbr

	switch updateValue := updateValue.(type) {
	case int:
		a.SetBalance(updateValue)
	case uuid.UUID:
		a.bankId = updateValue
	}

	return a

}
func (a *Account) DeleteAccount() *Account {

	a.SetIsActive()
	return a
}

func (a *Account) GetAccountNumber() uuid.UUID {
	return a.accountNumber
}
func (a *Account) GetBankId() uuid.UUID {
	return a.bankId
}
func (a *Account) GetCustomerId() uuid.UUID {
	return a.customerId
}
func (a *Account) GetIsActive() bool {
	return a.isActive
}
func (a *Account) SetIsActive() {
	a.isActive = false
}
func (a *Account) GetBalance() int {
	return a.balance
}
func (a *Account) SetBalance(balance int) {
	a.balance = balance
}

func (a *Account) DepositMoney(amount int) {
	a.balance += amount
	a.passbook.AddEntry(a.customerId, a.customerId, a.accountNumber, a.accountNumber, amount, "CREDIT")
}
func (a *Account) WithdrawMoney(amount int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if a.balance-amount >= 0 {
		a.balance -= amount
		a.passbook.AddEntry(a.customerId, a.customerId, a.accountNumber, a.accountNumber, amount, "")
	}
	panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())

}

func (a *Account) TransferMoney(receiver *Account, amount int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if a.balance-amount >= 0 {
		a.balance -= amount
		receiver.balance += amount
	}
	panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
}
