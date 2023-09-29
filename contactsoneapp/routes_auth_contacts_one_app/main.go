package main

import (
	"contactsoneapp/components/guru_contact_details"
	"contactsoneapp/components/guru_contacts"
	"contactsoneapp/components/guru_user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main Called....")
	handleMyRoutes()
}

func handleMyRoutes() {
	headerOK := handlers.AllowCredentials()
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})

	mainRouter := mux.NewRouter().StrictSlash(true)
	subRouter := mainRouter.PathPrefix("/api/v1/contacts-one-app").Subrouter()
	userRouter := guru_user.HandleRouter(subRouter)
	contactRouter := guru_contacts.HandleRouter(userRouter)
	guru_contact_details.HandleRouter(contactRouter)
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headerOK, originsOK, methodsOK)(mainRouter)))
}
