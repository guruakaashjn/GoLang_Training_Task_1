package guru_customer

import (
	"bankingapp/guru_account"
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
	return &Customer{
		customerId:   uuid.New(),
		firstName:    firstName,
		lastName:     lastName,
		totalBalance: 0,
		isAdmin:      isAdmin,
		isActive:     true,
		Accounts:     make([]*guru_account.Account, 0),
	}
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
func (c *Customer) CreateCustomer(firstName, lastName string) (flag bool, customer *Customer) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.isAdmin && c.isActive {
		return true, NewCustomer(firstName, lastName, false)
	}
	if c.isAdmin && !c.isActive {
		panic(guru_errors.NewAdminError(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}
func CreateAdmin(firstName, lastName string) *Customer {
	return NewCustomer(firstName, lastName, true)
}
func (c *Customer) AddAccount(bankId uuid.UUID, balance int) (flag bool, accountObject *guru_account.Account) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.isActive {
		accountObject = guru_account.CreateAccount(bankId, balance)

		c.Accounts = append(c.Accounts, accountObject)
		flag = true
		return flag, accountObject
	}
	panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())

}

func (c *Customer) ReadAllCustomers() []*Customer {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if c.isAdmin && c.isActive {
		return Customers
	}
	if c.isAdmin && !c.isActive {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
}

func (c *Customer) ReadCustomer(customer *Customer) (flag bool, returnCustomer *Customer) {
	flag = false
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)

		}
	}()
	if c.isAdmin && customer.isActive {
		return true, &Customer{
			customerId:   customer.customerId,
			firstName:    customer.firstName,
			lastName:     customer.lastName,
			totalBalance: customer.totalBalance,
			isAdmin:      customer.isAdmin,
			Accounts:     customer.Accounts,
		}
	}
	if c.isAdmin && !customer.isActive {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())
}
func (c *Customer) UpdateCustomer(customerId uuid.UUID, updateField, updateValue interface{}) (flag bool) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
			flag = false
		}
	}()
	if c.isAdmin && c.isActive {
		switch updateValue.(type) {
		case string:
			if updateField == "firstName" {
				c.firstName = updateValue.(string)
			} else {
				c.lastName = updateValue.(string)
			}

		}
		// switch updateField {
		// case "firstName":
		// 	c.firstName = updateValue
		// case "lastName":
		// 	c.lastName = updateValue
		// }

		flag = true
		panic(guru_errors.NewNotAUser(guru_errors.UpdatedUser).GetSpecificMessage())
	}
	if c.isAdmin && !c.isActive {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}
func (c *Customer) DeleteCustomer(customerId uuid.UUID) (flag bool) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
			flag = false
		}
	}()
	if c.isAdmin && c.isActive {

		for i := 0; i < len(Customers); i++ {
			if Customers[i].GetCustomerId() == customerId {
				for j := 0; j < len(Customers[i].Accounts); j++ {
					Customers[i].Accounts[j].DeleteAccount()
				}
				Customers[i].SetIsActve()
				break
			}
		}

		flag = true
		return flag

	}
	if c.isAdmin && !c.isActive {
		panic(guru_errors.NewNotAUser(guru_errors.DeletedUser).GetSpecificMessage())
	}
	panic(guru_errors.NewAdminError(guru_errors.NotAdmin).GetSpecificMessage())

}
