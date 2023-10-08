package controller

import (
	"bankingapp/components/guru_account/service"
	"bankingapp/components/log"
	"bankingapp/middleware/auth"
	useridtokenuseridverification "bankingapp/middleware/customerid_token_customerid_verification.go"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountController struct {
	log     log.Log
	service *service.AccountService
}

func NewAccountController(accountService *service.AccountService, log log.Log) *AccountController {
	return &AccountController{
		service: accountService,
		log:     log,
	}
}

func (controller *AccountController) RegisterRoutes(router *mux.Router) {
	accountRouter := router.PathPrefix("/customer/{customer-id}/account").Subrouter()
	accountRouter.HandleFunc("/", controller.CreateAccount).Methods(http.MethodPost)
	accountRouter.HandleFunc("/", controller.GetAllAccounts).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{id}", controller.UpdateAccount).Methods(http.MethodPut)
	accountRouter.HandleFunc("/{id}", controller.DeleteAccount).Methods(http.MethodDelete)
	accountRouter.HandleFunc("/{id}", controller.GetAccountById).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{id}/transfer-money", controller.TransferMoney).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id}/deposit-money", controller.DepositMoney).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id}/withdraw-money", controller.WithdrawMoney).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id}/passbook", controller.PassbookPrint).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{id}/get-available-offers", controller.GetAvailableOffers).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{id}/take-available-offer", controller.TakeAvailableOffer).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id}/chosen-offers", controller.ChosenOffers).Methods(http.MethodGet)

	accountRouter.Use(auth.IsUser)
	accountRouter.Use(useridtokenuseridverification.CheckUserVerify)

	fmt.Println("[Account register routes]")

}

// package guru_account

// import (
// 	account_controller "bankingapp/components/guru_account/controller"
// 	"bankingapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) {
// 	accountRouter := parentRouter.PathPrefix("/{customer-id}/account").Subrouter()

// 	guardedRouter := accountRouter.PathPrefix("/g").Subrouter()

// 	guardedRouter.HandleFunc("/", account_controller.CreateAccount).Methods("POST")
// 	guardedRouter.HandleFunc("/", account_controller.ReadAccountAll).Methods("GET")
// 	guardedRouter.HandleFunc("/{account-id}", account_controller.ReadAccountById).Methods("GET")
// 	guardedRouter.HandleFunc("/{account-id}", account_controller.UpdateAccountById).Methods("PUT")
// 	guardedRouter.HandleFunc("/{account-id}", account_controller.DeleteAccountById).Methods("DELETE")
// 	guardedRouter.HandleFunc("/{account-id}/deposit-money", account_controller.DepositMoney).Methods("POST")
// 	guardedRouter.HandleFunc("/{account-id}/withdraw-money", account_controller.WithdrawMoney).Methods("POST")
// 	guardedRouter.HandleFunc("/{account-id}/transfer-money", account_controller.TransferMoney).Methods("POST")
// 	guardedRouter.HandleFunc("/{account-id}/passbook", account_controller.PassbookPrint).Methods("POST")

// 	guardedRouter.Use(auth.IsUser)

// }
