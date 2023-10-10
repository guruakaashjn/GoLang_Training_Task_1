package controller

import (
	"bankingapp/errors"
	"bankingapp/models/customer"
	"bankingapp/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (controller *CustomerController) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	newAdmin := customer.Customer{}
	err := web.UnmarshalJSON(r, &newAdmin)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	newAdmin.IsAdmin = true
	err = controller.service.CreateCustomer(&newAdmin)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newAdmin)

}

func (controller *CustomerController) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	newCustomer := customer.Customer{}
	err := web.UnmarshalJSON(r, &newCustomer)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	newCustomer.IsAdmin = false
	err = controller.service.CreateCustomer(&newCustomer)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newCustomer)
}

func (controller *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	allCustomers := &[]customer.CustomerDTO{}
	var totalCount int
	limit, offset, err := web.ParseLimitAndOffset(r)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	givenAssociations := web.ParsePreloading(r)

	// columnNames, conditiond, operators, values := web.ParseForLike(r)

	allQueries, err := web.ParseQueryParams(r)
	if err != nil {
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	// var limit int
	// var offset int
	// var givenAssociations = make([]string, 0)
	// searchQueries := make(map[string]interface{}, 0)
	// for queryKey, queryValue := range allQueries {
	// 	if queryKey == "limit" {
	// 		limit = queryValue.(int)
	// 	}
	// 	if queryKey == "offset" {
	// 		offset = queryValue.(int)
	// 	}
	// 	if queryKey == "includes" {
	// 		givenAssociations = queryValue.([]string)
	// 	}
	// 	if queryKey != "limit" && queryKey != "offset" && queryKey != "includes" {
	// 		searchQueries[queryKey] = queryValue
	// 	}
	// }

	// println("--------------------------------------------------------------------", strings.Split(columnName.Encode(), "&")[0])

	// println("--------------------------------------------------------------------", strings.Split(columnName.Encode(), "=")[0])

	err = controller.service.GetAllCustomers(allCustomers, &totalCount, limit, offset, givenAssociations, allQueries)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	// allCustomersDTO := []customer.CustomerDTO{}
	// for _, customerObj := range *allCustomers {
	// 	customerDTOObj := customer.CustomerDTO{}
	// 	utils.ConvertUserObjectToDTOObject(&customerObj, &customerDTOObj)
	// 	allCustomersDTO = append(allCustomersDTO, customerDTOObj)
	// }
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allCustomers)
}

func (controller *CustomerController) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	requiredCustomer := customer.Customer{}

	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	givenAssociations := web.ParsePreloading(r)
	err = controller.service.GetCustomerById(&requiredCustomer, idTemp, givenAssociations)

	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredCustomer)
}

func (controller *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Customer to update")
	customerToUpdate := customer.Customer{}

	fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &customerToUpdate)
	if err != nil {
		fmt.Println("error from unmarshal JSON")
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)

	intID, err := strconv.Atoi(vars["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	customerToUpdate.ID = uint(intID)

	fmt.Println("Customer to update")
	fmt.Println(&customerToUpdate)
	err = controller.service.UpdateCustomer(&customerToUpdate)

	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, customerToUpdate)
}

func (controller *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	controller.log.Print("Delete customer call")
	customerToDelete := customer.Customer{}
	var err error
	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	customerToDelete.ID = uint(intId)
	err = controller.service.DeleteCustomer(&customerToDelete)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete customer successful.")
}

func (controller *CustomerController) TotalBalance(w http.ResponseWriter, r *http.Request) {

	requiredCustomer := &customer.Customer{}
	slugs := mux.Vars(r)

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	requiredCustomer.ID = uint(idTemp)
	var totalBalance int
	err = controller.service.TotalBalance(requiredCustomer, &totalBalance)
	// fmt.Println(totalBalance)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, totalBalance)

}

func (controller *CustomerController) AccountBalanceList(w http.ResponseWriter, r *http.Request) {
	requiredCustomer := &customer.Customer{}
	slugs := mux.Vars(r)

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	requiredCustomer.ID = uint(idTemp)

	var mapAccountBalance = make(map[uint]int, 0)
	err = controller.service.AccountBalanceList(requiredCustomer, mapAccountBalance)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	// fmt.Println(mapAccountBalance)
	web.RespondJSON(w, http.StatusOK, mapAccountBalance)

}

// package controller

// import (
// 	customer_service "bankingapp/components/guru_customer/service"
// 	"bankingapp/guru_errors"
// 	"bankingapp/middleware/auth"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// // var initialAdmin *customer_service.Customer = &customer_service.Customer{
// // 	CustomerId:   uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
// // 	FirstName:    "Admin Initial",
// // 	LastName:     "Admin surname",
// // 	TotalBalance: 0,
// // 	IsAdmin:      true,
// // 	IsActive:     true,
// // 	Accounts:     make([]*account_service.Account, 0),
// // }

// func GetAdminObjectFromCookie(w http.ResponseWriter, r *http.Request) (requiredAdmin *customer_service.Customer) {
// 	defer func() {
// 		if details := recover(); details != nil {
// 			fmt.Println(details)
// 		}
// 	}()

// 	cookie, err1 := r.Cookie("authone")
// 	if err1 != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(err1)
// 		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
// 	}

// 	token := cookie.Value
// 	payload, err2 := auth.Verify(token)
// 	if err2 != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(err2)
// 		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
// 	}

// 	requiredAdminTemp := customer_service.ReadCustomerByUserName(payload.UserName)
// 	if requiredAdminTemp == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DeletedAdmin)
// 	}

// 	requiredAdmin = requiredAdminTemp
// 	panic(guru_errors.NewUserError(guru_errors.AdminObjectFound).GetSpecificMessage())

// }

// func CreateAdmin(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside create admin controller function")
// 	var newCustomerTemp *customer_service.Customer
// 	err := json.NewDecoder(r.Body).Decode(&newCustomerTemp)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAdminError(guru_errors.CreateAdminFailed).GetSpecificMessage())
// 	}
// 	newCustomer := customer_service.CreateAdmin(
// 		newCustomerTemp.FirstName,
// 		newCustomerTemp.LastName,
// 		newCustomerTemp.UserName,
// 		newCustomerTemp.Password,
// 	)
// 	json.NewEncoder(w).Encode(newCustomer)
// 	// fmt.Println("Inside create admin controller function POST REQUEST DONE")
// 	panic(guru_errors.NewAdminError(guru_errors.CreateAdminSuccess).GetSpecificMessage())

// }

// func CreateCustomer(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside create customer controller function")
// 	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
// 	var newCustomerTemp *customer_service.Customer
// 	err := json.NewDecoder(r.Body).Decode(&newCustomerTemp)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewUserError(guru_errors.CreateCustomerFailed).GetSpecificMessage())
// 	}
// 	newCustomer := adminObject.CreateCustomer(
// 		newCustomerTemp.FirstName,
// 		newCustomerTemp.LastName,
// 		newCustomerTemp.UserName,
// 		newCustomerTemp.Password,
// 	)
// 	if newCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.CreateCustomerFailed)
// 		panic(guru_errors.NewUserError(guru_errors.CreateCustomerFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(newCustomer)
// 	// fmt.Println("Inside create customer controller function POST REQUEST DONE")
// 	panic(guru_errors.NewUserError(guru_errors.CreateCustomerSuccess).GetSpecificMessage())

// }

// func ReadCustomerAll(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside read all customers controller function")
// 	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
// 	requiredCustomers := adminObject.ReadAllCustomers()
// 	if requiredCustomers == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadCustomerFailed)
// 		panic(guru_errors.NewUserError(guru_errors.ReadCustomerFailed).GetSpecificMessage())
// 	}

// 	json.NewEncoder(w).Encode(requiredCustomers)
// 	// fmt.Println("Inside read all customers controller function GET REQUEST DONE")
// 	panic(guru_errors.NewUserError(guru_errors.ReadCustomerSuccess).GetSpecificMessage())

// }

// func ReadCustomerById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside read customer by id controller function")
// 	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
// 	slugs := mux.Vars(r)
// 	requiredCustomer := adminObject.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if requiredCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadCustomerFailed)
// 		panic(guru_errors.NewUserError(guru_errors.ReadCustomerFailed).GetSpecificMessage())
// 	}

// 	json.NewEncoder(w).Encode(requiredCustomer)

// 	// fmt.Println("Inside read customer by id controller function GET REQUEST DONE")
// 	panic(guru_errors.NewUserError(guru_errors.ReadCustomerSuccess).GetSpecificMessage())
// }

// func UpdateCustomerById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside update customer by id controller function")
// 	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
// 	var customerTempObject *customer_service.Customer
// 	err := json.NewDecoder(r.Body).Decode(&customerTempObject)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewUserError(guru_errors.UpdateCustomerFailed).GetSpecificMessage())
// 	}
// 	slugs := mux.Vars(r)
// 	// fmt.Println(slugs)
// 	updatedCustomer := adminObject.UpdateCustomerObject(uuid.MustParse(slugs["customer-id"]), customerTempObject)
// 	json.NewEncoder(w).Encode(updatedCustomer)
// 	// fmt.Println("Inside update customer by id controller function PUT REQUEST DONE")
// 	panic(guru_errors.NewUserError(guru_errors.UpdateCustomerSuccess).GetSpecificMessage())
// }

// func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside delete customer by id controller function")
// 	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
// 	slugs := mux.Vars(r)
// 	deletedCustomer := adminObject.DeleteCustomer(uuid.MustParse(slugs["customer-id"]))
// 	json.NewEncoder(w).Encode(deletedCustomer)

// 	// fmt.Println("Inside delete customer by id controller function DELETE REQUEST DONE")
// 	panic(guru_errors.NewUserError(guru_errors.DeleteCustomerSuccess).GetSpecificMessage())
// }

// func TotalBalance(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside total balance customer by id controller function")
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.TotalBalanceFailed)
// 		panic(guru_errors.NewUserError(guru_errors.TotalBalanceFailed).GetSpecificMessage())
// 	}

// 	currentCustomerTotalBalance := currentCustomer.GetTotalBalance()
// 	json.NewEncoder(w).Encode(currentCustomerTotalBalance)
// 	panic(guru_errors.NewUserError(guru_errors.TotalBalanceSuccess).GetSpecificMessage())
// }

// func AccountBalanceList(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside account balance list customer by id controller function")
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.AccountBalanceListFailed)
// 		panic(guru_errors.NewUserError(guru_errors.AccountBalanceListFailed).GetSpecificMessage())
// 	}
// 	currentCustomerAccountBalanceList := currentCustomer.GetAllIndividualAccountBalance()
// 	json.NewEncoder(w).Encode(currentCustomerAccountBalanceList)
// 	panic(guru_errors.NewUserError(guru_errors.AccountBalanceListSuccess).GetSpecificMessage())
// }
