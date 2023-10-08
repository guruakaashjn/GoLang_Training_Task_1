package controller

import (
	"bankingapp/components/guru_bank/service"
	"bankingapp/components/log"
	"bankingapp/middleware/auth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BankController struct {
	log     log.Log
	service *service.BankService
}

func NewBankController(bankService *service.BankService, log log.Log) *BankController {
	return &BankController{
		service: bankService,
		log:     log,
	}
}

func (controller *BankController) RegisterRoutes(router *mux.Router) {
	bankRouter := router.PathPrefix("/bank").Subrouter()
	bankRouter.HandleFunc("/", controller.RegisterBank).Methods(http.MethodPost)
	bankRouter.HandleFunc("/", controller.GetAllBanks).Methods(http.MethodGet)

	bankRouter.HandleFunc("/all-bank-networth", controller.AllBankNetWorth).Methods(http.MethodGet)
	bankRouter.HandleFunc("/all-bank-balance-map", controller.AllBankBalanceMap).Methods("GET")

	bankRouter.HandleFunc("/{id}", controller.UpdateBank).Methods(http.MethodPut)
	bankRouter.HandleFunc("/{id}", controller.DeleteBank).Methods(http.MethodDelete)
	bankRouter.HandleFunc("/{id}", controller.GetBankById).Methods(http.MethodGet)

	bankRouter.HandleFunc("/{id}/bank-networth", controller.BankNetworth).Methods(http.MethodGet)
	bankRouter.HandleFunc("/{id}/bank-balance-map", controller.BankBalanceMap).Methods(http.MethodGet)
	bankRouter.Use(auth.IsAdmin)

	fmt.Println("[Bank register routes]")
}

// package guru_bank

// import (
// 	bank_controller "bankingapp/components/guru_bank/controller"
// 	"bankingapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) {
// 	bankRouter := parentRouter.PathPrefix("/bank").Subrouter()

// 	guardedRouter := bankRouter.PathPrefix("/g").Subrouter()

// 	guardedRouter.HandleFunc("/", bank_controller.CreateBank).Methods("POST")
// 	guardedRouter.HandleFunc("/", bank_controller.ReadBankAll).Methods("GET")
// 	guardedRouter.HandleFunc("/all-bank-networth", bank_controller.NetWorthEachBank).Methods("GET")
// 	guardedRouter.HandleFunc("/all-bank-balance-map", bank_controller.BankNameBalanceMapAll).Methods("POST")

// 	guardedRouter.HandleFunc("/{bank-id}", bank_controller.ReadBankById).Methods("GET")
// 	guardedRouter.HandleFunc("/{bank-id}", bank_controller.UpdateBankById).Methods("PUT")
// 	guardedRouter.HandleFunc("/{bank-id}", bank_controller.DeleteBankById).Methods("DELETE")

// 	guardedRouter.HandleFunc("/{bank-id}/networth-given-bank", bank_controller.NetWorthGivenBank).Methods("GET")
// 	guardedRouter.HandleFunc("/{bank-id}/bank-balance-map", bank_controller.BankNameBalanceMapByBankId).Methods("POST")

// 	guardedRouter.Use(auth.IsAdmin)

// }
