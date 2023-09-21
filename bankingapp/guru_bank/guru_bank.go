package guru_bank

import (
	"bankingapp/guru_account"
	"bankingapp/guru_customer"
	"bankingapp/guru_errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

var Banks = make([]*Bank, 0)

type Bank struct {
	bankId       uuid.UUID
	fullName     string
	abbreviation string
	isActive     bool
	Accounts     []*guru_account.Account
}

func NewBank(fullName string) *Bank {
	var abbr string = setAbbreviation(fullName)
	return &Bank{
		bankId:       uuid.New(),
		fullName:     fullName,
		abbreviation: abbr,
		isActive:     true,
		Accounts:     make([]*guru_account.Account, 0),
	}
}

func (b *Bank) GetAbbreviation() string {
	return b.abbreviation
}

func setAbbreviation(fullName string) (abbr string) {
	abbr = fullName[0:4] + fullName[len(fullName)-4:]
	max := rand.Intn(10000-5000) + 5000
	min := rand.Intn(4999-0) + 0
	rnd := rand.Intn(max-min) + min
	abbr += strconv.Itoa(rnd)

	if checkAbbreviation(abbr) {
		return abbr
	}

	return setAbbreviation(fullName)
}

func checkAbbreviation(abbr string) (flag bool) {
	for i := 0; i < len(Banks); i++ {
		if Banks[i].GetAbbreviation() == abbr {
			return false
		}

	}
	return true
}

func CreateBank(c *guru_customer.Customer, fullName string) (flag bool, bankName *Bank) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.GetIsAdmin() && c.GetIsActive() {
		flag = true
		return true, NewBank(fullName)

	}
	if c.GetIsAdmin() && !c.GetIsActive() {
		panic(guru_errors.NewAdminError(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
}

func (b *Bank) AddAccount(c *guru_customer.Customer, a *guru_account.Account) (flag bool) {

	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.GetIsActive() {
		c.Accounts = append(c.Accounts, a)
		flag = true
		return flag
	}
	panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())

}
func (b *Bank) ReadBank(c *guru_customer.Customer) (flag bool, bank *Bank) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()

	if c.GetIsAdmin() && c.GetIsActive() {
		flag = true
		return true, &Bank{
			bankId:       b.bankId,
			fullName:     b.fullName,
			abbreviation: b.abbreviation,
			isActive:     b.isActive,
			Accounts:     b.Accounts,
		}
	}
	if c.GetIsAdmin() && !c.GetIsActive() {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}
func (b *Bank) ReadAllBanks(c *guru_customer.Customer) []*Bank {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.GetIsAdmin() && c.GetIsActive() {
		return Banks
	}
	if c.GetIsAdmin() && !c.GetIsAdmin() {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
}

func (b *Bank) UpdateBank(c *guru_customer.Customer, updateValue string) (flag bool) {
	flag = false

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.GetIsAdmin() && c.GetIsActive() {

		b.fullName = updateValue
		b.abbreviation = setAbbreviation(b.fullName)
		flag = true
	}
	if c.GetIsAdmin() && !c.GetIsAdmin() {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}

func (b *Bank) DeleteBank(c *guru_customer.Customer) (flag bool) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()

	if c.GetIsAdmin() && c.GetIsActive() && len(b.Accounts) == 0 {
		b.isActive = false
		flag = true
		panic(guru_errors.NewBankError(guru_errors.DeletedBank).GetSpecificMessage())
	}
	if c.GetIsAdmin() && c.GetIsActive() && len(b.Accounts) != 0 {
		panic(guru_errors.NewBankError(guru_errors.BankContainsAccounts).GetSpecificMessage())
	}
	if c.GetIsAdmin() && !c.GetIsActive() {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}
