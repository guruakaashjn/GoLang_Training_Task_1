package guru_customer

import (
	customer_controller "bankingapp/components/guru_customer/controller"
	"bankingapp/middleware/auth"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) *mux.Router {

	parentRouter.HandleFunc("/admin", customer_controller.CreateAdmin).Methods("POST")
	parentRouter.HandleFunc("/login", customer_controller.Login).Methods("POST")

	customerRouter := parentRouter.PathPrefix("/customer").Subrouter()
	guardedRouter := customerRouter.PathPrefix("/g").Subrouter()

	guardedRouter.HandleFunc("/", customer_controller.CreateCustomer).Methods("POST")
	guardedRouter.HandleFunc("/", customer_controller.ReadCustomerAll).Methods("GET")
	guardedRouter.HandleFunc("/{customer-id}", customer_controller.ReadCustomerById).Methods("GET")
	guardedRouter.HandleFunc("/{customer-id}", customer_controller.UpdateCustomerById).Methods("PUT")
	guardedRouter.HandleFunc("/{customer-id}", customer_controller.DeleteCustomerById).Methods("DELETE")
	guardedRouter.HandleFunc("/{customer-id}/total-balance", customer_controller.TotalBalance).Methods("GET")
	guardedRouter.HandleFunc("/{customer-id}/account-balance-list", customer_controller.AccountBalanceList).Methods("GET")

	guardedRouter.Use(auth.IsAdmin)
	return customerRouter
}
