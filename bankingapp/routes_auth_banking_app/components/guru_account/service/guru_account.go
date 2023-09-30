package service

import (
	"bankingapp/components/guru_passbook"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountNumber uuid.UUID
	BankId        uuid.UUID
	CustomerId    uuid.UUID
	IsActive      bool
	Balance       int
	Passbook      *guru_passbook.Passbook
}

func NewAccount(BankId uuid.UUID, CustomerId uuid.UUID, Balance int) *Account {
	AccountNumber := uuid.New()
	var newPassbook *guru_passbook.Passbook = guru_passbook.NewPassbook(CustomerId, AccountNumber, Balance)
	fmt.Println("Passbook from NewAccount() guru_account.go service : ", newPassbook)
	var newAccountObject *Account = &Account{
		AccountNumber: AccountNumber,
		BankId:        BankId,
		CustomerId:    CustomerId,
		IsActive:      true,
		Balance:       Balance,
		Passbook:      newPassbook,
	}
	fmt.Println(newAccountObject.Passbook)

	return newAccountObject
}

func CreateAccount(BankId uuid.UUID, CustomerId uuid.UUID, Balance int) (account *Account) {

	return NewAccount(BankId, CustomerId, Balance)

}

func (a *Account) ReadAccount() (bool, *Account) {
	if a.IsActive {
		return true, a
	}
	return false, a

}
func (a *Account) UpdateAccount(updateField string, updateValue interface{}) *Account {
	// Balance, abbr

	switch updateValue := updateValue.(type) {
	case int:
		a.SetBalance(updateValue)
	case uuid.UUID:
		a.BankId = updateValue
	}

	return a

}
func (a *Account) DeleteAccount() *Account {

	a.SetIsActive()
	return a
}

func (a *Account) GetAccountNumber() uuid.UUID {
	return a.AccountNumber
}
func (a *Account) GetBankId() uuid.UUID {
	return a.BankId
}
func (a *Account) GetCustomerId() uuid.UUID {
	return a.CustomerId
}
func (a *Account) GetIsActive() bool {
	return a.IsActive
}
func (a *Account) SetIsActive() {
	a.IsActive = false
}
func (a *Account) GetBalance() int {
	return a.Balance
}
func (a *Account) SetBalance(Balance int) {
	a.Balance = Balance
}

func (a *Account) GetPassbook(startDate, endDate time.Time) *guru_passbook.Passbook {
	// fmt.Println("Get Passbook Func : ", startDate, "    ", endDate)
	return a.Passbook.ReadPassbook(startDate, endDate)
}
func (a *Account) DepositMoney(amount int) *Account {
	a.Balance += amount
	a.Passbook.AddEntry(a.CustomerId, a.CustomerId, a.AccountNumber, a.AccountNumber, amount, "CREDIT")
	return a
}

func (a *Account) WithdrawMoney(amount int) *Account {
	a.Balance -= amount
	a.Passbook.AddEntry(a.CustomerId, a.CustomerId, a.AccountNumber, a.AccountNumber, amount, "DEBIT")
	return a

}

func (a *Account) TransferMoney(receiver *Account, amount int) (*Account, *Account) {

	a.Balance -= amount
	receiver.Balance += amount
	a.Passbook.AddEntry(a.CustomerId, receiver.CustomerId, a.AccountNumber, receiver.AccountNumber, amount, "DEBIT")
	receiver.Passbook.AddEntry(a.CustomerId, receiver.CustomerId, a.AccountNumber, receiver.AccountNumber, amount, "CREDIT")
	return a, receiver
}
