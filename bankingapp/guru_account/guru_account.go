package guru_account

import (
	"bankingapp/guru_errors"
	"fmt"

	"github.com/google/uuid"
)

type Account struct {
	accountNumber uuid.UUID
	bank          uuid.UUID
	isActive      bool
	balance       int
}

func NewAccount(bankId uuid.UUID, balance int) *Account {
	return &Account{
		accountNumber: uuid.New(),
		bank:          uuid.New(),
		isActive:      true,
		balance:       balance,
	}
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

func CreateAccount(bankId uuid.UUID, balance int) *Account {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if balance >= 1000 {
		return NewAccount(bankId, balance)
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
			bank:          a.bank,
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
			a.bank = updateValue.(uuid.UUID)
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
