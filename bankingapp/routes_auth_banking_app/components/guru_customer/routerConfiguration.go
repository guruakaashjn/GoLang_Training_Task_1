package guru_customer

import (
	customer_controller "bankingapp/components/guru_customer/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) *mux.Router {

	parentRouter.HandleFunc("/admin", customer_controller.CreateAdmin).Methods("POST")
	customerRouter := parentRouter.PathPrefix("/customer").Subrouter()
	customerRouter.HandleFunc("/", customer_controller.CreateCustomer).Methods("POST")
	customerRouter.HandleFunc("/", customer_controller.ReadCustomerAll).Methods("GET")
	customerRouter.HandleFunc("/{customer-id}", customer_controller.ReadCustomerById).Methods("GET")
	customerRouter.HandleFunc("/{customer-id}", customer_controller.UpdateCustomerById).Methods("PUT")
	customerRouter.HandleFunc("/{customer-id}", customer_controller.DeleteCustomerById).Methods("DELETE")
	customerRouter.HandleFunc("/{customer-id}/total-balance", customer_controller.TotalBalance).Methods("GET")
	customerRouter.HandleFunc("/{customer-id}/account-balance-list", customer_controller.AccountBalanceList).Methods("GET")
	return customerRouter
}
