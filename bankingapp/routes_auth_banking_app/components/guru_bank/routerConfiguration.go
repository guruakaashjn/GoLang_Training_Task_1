package guru_bank

import (
	bank_controller "bankingapp/components/guru_bank/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	bankRouter := parentRouter.PathPrefix("/bank").Subrouter()
	bankRouter.HandleFunc("/", bank_controller.CreateBank).Methods("POST")
	bankRouter.HandleFunc("/", bank_controller.ReadBankAll).Methods("GET")
	bankRouter.HandleFunc("/all-bank-networth", bank_controller.NetWorthEachBank).Methods("GET")
	bankRouter.HandleFunc("/all-bank-balance-map", bank_controller.BankNameBalanceMapAll).Methods("POST")

	bankRouter.HandleFunc("/{bank-id}", bank_controller.ReadBankById).Methods("GET")
	bankRouter.HandleFunc("/{bank-id}", bank_controller.UpdateBankById).Methods("PUT")
	bankRouter.HandleFunc("/{bank-id}", bank_controller.DeleteBankById).Methods("DELETE")

	bankRouter.HandleFunc("/{bank-id}/networth-given-bank", bank_controller.NetWorthGivenBank).Methods("GET")
	bankRouter.HandleFunc("/{bank-id}/bank-balance-map", bank_controller.BankNameBalanceMapByBankId).Methods("POST")

}
