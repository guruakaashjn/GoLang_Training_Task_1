package guru_account

import (
	account_controller "bankingapp/components/guru_account/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	accountRouter := parentRouter.PathPrefix("/{customer-id}/account").Subrouter()
	accountRouter.HandleFunc("/", account_controller.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/", account_controller.ReadAccountAll).Methods("GET")
	accountRouter.HandleFunc("/{account-id}", account_controller.ReadAccountById).Methods("GET")
	accountRouter.HandleFunc("/{account-id}", account_controller.UpdateAccountById).Methods("PUT")
	accountRouter.HandleFunc("/{account-id}", account_controller.DeleteAccountById).Methods("DELETE")
	accountRouter.HandleFunc("/{account-id}/deposit-money", account_controller.DepositMoney).Methods("POST")
	accountRouter.HandleFunc("/{account-id}/withdraw-money", account_controller.WithdrawMoney).Methods("POST")
	accountRouter.HandleFunc("/{account-id}/transfer-money", account_controller.TransferMoney).Methods("POST")
	accountRouter.HandleFunc("/{account-id}/passbook", account_controller.PassbookPrint).Methods("POST")

}
