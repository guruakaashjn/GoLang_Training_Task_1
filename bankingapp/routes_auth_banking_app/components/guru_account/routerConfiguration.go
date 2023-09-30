package guru_account

import (
	account_controller "bankingapp/components/guru_account/controller"
	"bankingapp/middleware/auth"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	accountRouter := parentRouter.PathPrefix("/{customer-id}/account").Subrouter()

	guardedRouter := accountRouter.PathPrefix("/g").Subrouter()

	guardedRouter.HandleFunc("/", account_controller.CreateAccount).Methods("POST")
	guardedRouter.HandleFunc("/", account_controller.ReadAccountAll).Methods("GET")
	guardedRouter.HandleFunc("/{account-id}", account_controller.ReadAccountById).Methods("GET")
	guardedRouter.HandleFunc("/{account-id}", account_controller.UpdateAccountById).Methods("PUT")
	guardedRouter.HandleFunc("/{account-id}", account_controller.DeleteAccountById).Methods("DELETE")
	guardedRouter.HandleFunc("/{account-id}/deposit-money", account_controller.DepositMoney).Methods("POST")
	guardedRouter.HandleFunc("/{account-id}/withdraw-money", account_controller.WithdrawMoney).Methods("POST")
	guardedRouter.HandleFunc("/{account-id}/transfer-money", account_controller.TransferMoney).Methods("POST")
	guardedRouter.HandleFunc("/{account-id}/passbook", account_controller.PassbookPrint).Methods("POST")

	guardedRouter.Use(auth.IsUser)

}
