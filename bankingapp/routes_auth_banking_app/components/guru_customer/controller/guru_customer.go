package controller

import (
	account_service "bankingapp/components/guru_account/service"
	customer_service "bankingapp/components/guru_customer/service"
	"bankingapp/guru_errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var initialAdmin *customer_service.Customer = &customer_service.Customer{
	CustomerId:   uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
	FirstName:    "Admin Initial",
	LastName:     "Admin surname",
	TotalBalance: 0,
	IsAdmin:      true,
	IsActive:     true,
	Accounts:     make([]*account_service.Account, 0),
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside create admin controller function")
	var newCustomerTemp *customer_service.Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomerTemp)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAdminError(guru_errors.CreateAdminFailed).GetSpecificMessage())
	}
	newCustomer := customer_service.CreateAdmin(
		newCustomerTemp.FirstName,
		newCustomerTemp.LastName,
		newCustomerTemp.UserName,
		newCustomerTemp.Password,
	)
	json.NewEncoder(w).Encode(newCustomer)
	// fmt.Println("Inside create admin controller function POST REQUEST DONE")
	panic(guru_errors.NewAdminError(guru_errors.CreateAdminSuccess).GetSpecificMessage())

}
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside create customer controller function")
	var newCustomerTemp *customer_service.Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomerTemp)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.CreateCustomerFailed).GetSpecificMessage())
	}
	newCustomer := initialAdmin.CreateCustomer(
		newCustomerTemp.FirstName,
		newCustomerTemp.LastName,
		newCustomerTemp.UserName,
		newCustomerTemp.Password,
	)
	if newCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.CreateCustomerFailed)
		panic(guru_errors.NewUserError(guru_errors.CreateCustomerFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(newCustomer)
	// fmt.Println("Inside create customer controller function POST REQUEST DONE")
	panic(guru_errors.NewUserError(guru_errors.CreateCustomerSuccess).GetSpecificMessage())

}

func ReadCustomerAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read all customers controller function")
	requiredCustomers := initialAdmin.ReadAllCustomers()
	if requiredCustomers == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadCustomerFailed)
		panic(guru_errors.NewUserError(guru_errors.ReadCustomerFailed).GetSpecificMessage())
	}

	json.NewEncoder(w).Encode(requiredCustomers)
	// fmt.Println("Inside read all customers controller function GET REQUEST DONE")
	panic(guru_errors.NewUserError(guru_errors.ReadCustomerSuccess).GetSpecificMessage())

}
func ReadCustomerById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read customer by id controller function")
	slugs := mux.Vars(r)
	requiredCustomer := initialAdmin.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if requiredCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadCustomerFailed)
		panic(guru_errors.NewUserError(guru_errors.ReadCustomerFailed).GetSpecificMessage())
	}

	json.NewEncoder(w).Encode(requiredCustomer)

	// fmt.Println("Inside read customer by id controller function GET REQUEST DONE")
	panic(guru_errors.NewUserError(guru_errors.ReadCustomerSuccess).GetSpecificMessage())
}
func UpdateCustomerById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside update customer by id controller function")
	var customerTempObject *customer_service.Customer
	err := json.NewDecoder(r.Body).Decode(&customerTempObject)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.UpdateCustomerFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	// fmt.Println(slugs)
	updatedCustomer := initialAdmin.UpdateCustomerObject(uuid.MustParse(slugs["customer-id"]), customerTempObject)
	json.NewEncoder(w).Encode(updatedCustomer)
	// fmt.Println("Inside update customer by id controller function PUT REQUEST DONE")
	panic(guru_errors.NewUserError(guru_errors.UpdateCustomerSuccess).GetSpecificMessage())
}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside delete customer by id controller function")
	slugs := mux.Vars(r)
	deletedCustomer := initialAdmin.DeleteCustomer(uuid.MustParse(slugs["customer-id"]))
	json.NewEncoder(w).Encode(deletedCustomer)

	// fmt.Println("Inside delete customer by id controller function DELETE REQUEST DONE")
	panic(guru_errors.NewUserError(guru_errors.DeleteCustomerSuccess).GetSpecificMessage())
}

func TotalBalance(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside total balance customer by id controller function")
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.TotalBalanceFailed)
		panic(guru_errors.NewUserError(guru_errors.TotalBalanceFailed).GetSpecificMessage())
	}

	currentCustomerTotalBalance := currentCustomer.GetTotalBalance()
	json.NewEncoder(w).Encode(currentCustomerTotalBalance)
	panic(guru_errors.NewUserError(guru_errors.TotalBalanceSuccess).GetSpecificMessage())
}
func AccountBalanceList(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside account balance list customer by id controller function")
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.AccountBalanceListFailed)
		panic(guru_errors.NewUserError(guru_errors.AccountBalanceListFailed).GetSpecificMessage())
	}
	currentCustomerAccountBalanceList := currentCustomer.GetAllIndividualAccountBalance()
	json.NewEncoder(w).Encode(currentCustomerAccountBalanceList)
	panic(guru_errors.NewUserError(guru_errors.AccountBalanceListSuccess).GetSpecificMessage())
}
