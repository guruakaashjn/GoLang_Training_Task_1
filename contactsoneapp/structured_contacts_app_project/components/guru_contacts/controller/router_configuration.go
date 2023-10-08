package controller

import (
	"contactsoneapp/components/guru_contacts/service"
	"contactsoneapp/components/log"
	"contactsoneapp/middleware/auth"
	useridtokenuseridverification "contactsoneapp/middleware/userid_token_userid_verification"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ContactController struct {
	log     log.Log
	service *service.ContactService
}

func NewContactController(contactService *service.ContactService, log log.Log) *ContactController {
	return &ContactController{
		service: contactService,
		log:     log,
	}
}

func (controller *ContactController) RegisterRoutes(router *mux.Router) {
	contactRouter := router.PathPrefix("/user/{user-id}/contact").Subrouter()
	contactRouter.HandleFunc("/", controller.CreateContact).Methods(http.MethodPost)
	contactRouter.HandleFunc("/", controller.GetAllContacts).Methods(http.MethodGet)
	contactRouter.HandleFunc("/{id}", controller.GetContactById).Methods(http.MethodGet)
	contactRouter.HandleFunc("/{id}", controller.UpdateContact).Methods(http.MethodPut)
	contactRouter.HandleFunc("/{id}", controller.DeleteContact).Methods(http.MethodDelete)
	contactRouter.Use(auth.IsUser)
	contactRouter.Use(useridtokenuseridverification.CheckUserVerify)
	fmt.Println("[Contact register routes]")
}

// package guru_contacts

// package guru_contacts

// import (
// 	devicecontroller "contactsoneapp/components/guru_contacts/controller"
// 	"contactsoneapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) *mux.Router {
// 	contactsRouter := parentRouter.PathPrefix("/{user-id}/contacts").Subrouter()
// 	guidedRouter := contactsRouter.PathPrefix("/g").Subrouter()
// 	guidedRouter.HandleFunc("/", devicecontroller.ReadContactsAll).Methods("GET")
// 	guidedRouter.HandleFunc("/", devicecontroller.CreateContact).Methods("POST")
// 	guidedRouter.HandleFunc("/{contact-id}", devicecontroller.ReadContactById).Methods("GET")
// 	guidedRouter.HandleFunc("/{contact-id}", devicecontroller.UpdateContactById).Methods("PUT")
// 	guidedRouter.HandleFunc("/{contact-id}", devicecontroller.DeleteContactById).Methods("DELETE")
// 	guidedRouter.Use(auth.IsUser)
// 	return contactsRouter

// }
