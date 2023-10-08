package controller

import (
	"bankingapp/components/guru_customer/service"
	"bankingapp/components/log"
	"bankingapp/middleware/auth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	log     log.Log
	service *service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService, log log.Log) *CustomerController {
	return &CustomerController{
		service: customerService,
		log:     log,
	}
}

func (controller *CustomerController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin-priv", controller.RegisterAdmin).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	customerRouter := router.PathPrefix("/customer").Subrouter()
	customerRouter.HandleFunc("/", controller.RegisterCustomer).Methods(http.MethodPost)
	customerRouter.HandleFunc("/", controller.GetAllCustomers).Methods(http.MethodGet)
	customerRouter.HandleFunc("/{id}", controller.UpdateCustomer).Methods(http.MethodPut)
	customerRouter.HandleFunc("/{id}", controller.DeleteCustomer).Methods(http.MethodDelete)
	customerRouter.HandleFunc("/{id}", controller.GetCustomerById).Methods(http.MethodGet)
	customerRouter.HandleFunc("/{id}/total-balance", controller.TotalBalance).Methods(http.MethodGet)
	customerRouter.HandleFunc("/{id}/account-balance-list", controller.AccountBalanceList).Methods(http.MethodGet)
	customerRouter.Use(auth.IsAdmin)

	fmt.Println("[Customer register routes]")
}

// package guru_customer

// import (
// 	customer_controller "bankingapp/components/guru_customer/controller"
// 	"bankingapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) *mux.Router {

// 	parentRouter.HandleFunc("/admin", customer_controller.CreateAdmin).Methods("POST")
// 	parentRouter.HandleFunc("/login", customer_controller.Login).Methods("POST")

// 	customerRouter := parentRouter.PathPrefix("/customer").Subrouter()
// 	guardedRouter := customerRouter.PathPrefix("/g").Subrouter()

// 	guardedRouter.HandleFunc("/", customer_controller.CreateCustomer).Methods("POST")
// 	guardedRouter.HandleFunc("/", customer_controller.ReadCustomerAll).Methods("GET")
// 	guardedRouter.HandleFunc("/{customer-id}", customer_controller.ReadCustomerById).Methods("GET")
// 	guardedRouter.HandleFunc("/{customer-id}", customer_controller.UpdateCustomerById).Methods("PUT")
// 	guardedRouter.HandleFunc("/{customer-id}", customer_controller.DeleteCustomerById).Methods("DELETE")
// 	guardedRouter.HandleFunc("/{customer-id}/total-balance", customer_controller.TotalBalance).Methods("GET")
// 	guardedRouter.HandleFunc("/{customer-id}/account-balance-list", customer_controller.AccountBalanceList).Methods("GET")

// 	guardedRouter.Use(auth.IsAdmin)
// 	return customerRouter
// }
