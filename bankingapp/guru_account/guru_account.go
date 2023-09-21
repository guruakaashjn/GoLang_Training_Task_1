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
func (a *Account) GetAccountNumber() uuid.UUID {
	return a.accountNumber
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
		if a := recover(); a != nil {
			fmt.Println(a)
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
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if a.balance-amount >= 0 {
		a.balance -= amount
		receiver.balance += amount
	}
	panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())
}

func CreateAccount(bankId uuid.UUID, customerId uuid.UUID, balance int) *Account {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if balance >= 1000 {
		return NewAccount(bankId, customerId, balance)
	}
	panic(guru_errors.NewAccountError(guru_errors.InSufficientBalance).GetSpecificMessage())

}

func (a *Account) ReadAccount() *Account {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if a.isActive {
		return &Account{
			accountNumber: a.accountNumber,
			bankId:        a.bankId,
			balance:       a.balance,
		}
	}
	panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())

}
func (a *Account) UpdateAccount(updateField string, updateValue interface{}) (flag bool) {
	// balance, abbr
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()
	if a.isActive {
		switch updateValue.(type) {
		case int:
			a.SetBalance(updateValue.(int))
		case uuid.UUID:
			a.bankId = updateValue.(uuid.UUID)
		}

		flag = true
		panic(guru_errors.NewAccountError(guru_errors.UpdatedAccount).GetSpecificMessage())
	}
	panic(guru_errors.NewAccountError(guru_errors.DeletedAccountStatus).GetSpecificMessage())

}
func (a *Account) DeleteAccount() (flag bool) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()
	if a.isActive {
		a.SetIsActive()
		flag = true
		panic(guru_errors.NewAccountError(guru_errors.DeletedAccount).GetSpecificMessage())
	}
	panic(guru_errors.NewAccountError(guru_errors.DeletedAccountAlready).GetSpecificMessage())
}
