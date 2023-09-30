package guru_bank

import (
	bank_controller "bankingapp/components/guru_bank/controller"
	"bankingapp/middleware/auth"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	bankRouter := parentRouter.PathPrefix("/bank").Subrouter()

	guardedRouter := bankRouter.PathPrefix("/g").Subrouter()

	guardedRouter.HandleFunc("/", bank_controller.CreateBank).Methods("POST")
	guardedRouter.HandleFunc("/", bank_controller.ReadBankAll).Methods("GET")
	guardedRouter.HandleFunc("/all-bank-networth", bank_controller.NetWorthEachBank).Methods("GET")
	guardedRouter.HandleFunc("/all-bank-balance-map", bank_controller.BankNameBalanceMapAll).Methods("POST")

	guardedRouter.HandleFunc("/{bank-id}", bank_controller.ReadBankById).Methods("GET")
	guardedRouter.HandleFunc("/{bank-id}", bank_controller.UpdateBankById).Methods("PUT")
	guardedRouter.HandleFunc("/{bank-id}", bank_controller.DeleteBankById).Methods("DELETE")

	guardedRouter.HandleFunc("/{bank-id}/networth-given-bank", bank_controller.NetWorthGivenBank).Methods("GET")
	guardedRouter.HandleFunc("/{bank-id}/bank-balance-map", bank_controller.BankNameBalanceMapByBankId).Methods("POST")

	guardedRouter.Use(auth.IsAdmin)

}
