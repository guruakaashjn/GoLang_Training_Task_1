package main

import (
	"bankingapp/components/guru_account"
	"bankingapp/components/guru_bank"
	guru_customer "bankingapp/components/guru_customer"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main Called....")
	// var abbr string = guru_bank.GetAbbreviation("Bank of India")
	// fmt.Println(abbr)
	HandleMyRoutes()
}

func HandleMyRoutes() {
	headerOK := handlers.AllowCredentials()
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"Get", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})

	mainRouter := mux.NewRouter().StrictSlash(true)
	subRouter := mainRouter.PathPrefix("/api/v1/banking-app").Subrouter()
	customerRouter := guru_customer.HandleRouter(subRouter)
	guru_bank.HandleRouter(subRouter)
	guru_account.HandleRouter(customerRouter)

	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headerOK, originsOK, methodsOK)(mainRouter)))
}
