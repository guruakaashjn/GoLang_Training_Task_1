package controller

import (
	customer_service "bankingapp/components/guru_customer/service"
	"bankingapp/guru_errors"
	"bankingapp/middleware/auth"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// var initialAdmin *customer_service.Customer = &customer_service.Customer{
// 	CustomerId:   uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
// 	FirstName:    "Admin Initial",
// 	LastName:     "Admin surname",
// 	TotalBalance: 0,
// 	IsAdmin:      true,
// 	IsActive:     true,
// 	Accounts:     make([]*account_service.Account, 0),
// }

func GetAdminObjectFromCookie(w http.ResponseWriter, r *http.Request) (requiredAdmin *customer_service.Customer) {
	defer func() {
		if details := recover(); details != nil {
			fmt.Println(details)
		}
	}()

	cookie, err1 := r.Cookie("authone")
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err1)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}

	token := cookie.Value
	payload, err2 := auth.Verify(token)
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err2)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}

	requiredAdminTemp := customer_service.ReadCustomerByUserName(payload.UserName)
	if requiredAdminTemp == nil {
		json.NewEncoder(w).Encode(guru_errors.DeletedAdmin)
	}

	requiredAdmin = requiredAdminTemp
	panic(guru_errors.NewUserError(guru_errors.AdminObjectFound).GetSpecificMessage())

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
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var newCustomerTemp *customer_service.Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomerTemp)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.CreateCustomerFailed).GetSpecificMessage())
	}
	newCustomer := adminObject.CreateCustomer(
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
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	requiredCustomers := adminObject.ReadAllCustomers()
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
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	requiredCustomer := adminObject.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
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
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var customerTempObject *customer_service.Customer
	err := json.NewDecoder(r.Body).Decode(&customerTempObject)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.UpdateCustomerFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	// fmt.Println(slugs)
	updatedCustomer := adminObject.UpdateCustomerObject(uuid.MustParse(slugs["customer-id"]), customerTempObject)
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
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	deletedCustomer := adminObject.DeleteCustomer(uuid.MustParse(slugs["customer-id"]))
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
