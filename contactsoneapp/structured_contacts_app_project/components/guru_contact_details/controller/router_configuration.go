package controller

import (
	"contactsoneapp/components/guru_contact_details/service"
	"contactsoneapp/components/log"
	"contactsoneapp/middleware/auth"
	useridtokenuseridverification "contactsoneapp/middleware/userid_token_userid_verification"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ContactInfoController struct {
	log     log.Log
	service service.ContactDetailsService
}

func NewContactInfoController(contactDetailsService *service.ContactDetailsService, log log.Log) *ContactInfoController {
	return &ContactInfoController{
		service: *contactDetailsService,
		log:     log,
	}
}
func (controller *ContactInfoController) RegisterRoutes(router *mux.Router) {
	contactDetailsRouter := router.PathPrefix("/user/{user-id}/contact/{contact-id}/contact-details").Subrouter()
	contactDetailsRouter.HandleFunc("/", controller.CreateContactDetails).Methods(http.MethodPost)
	contactDetailsRouter.HandleFunc("/", controller.GetAllContactDetails).Methods(http.MethodGet)
	contactDetailsRouter.HandleFunc("/{id}", controller.GetContactDetailsById).Methods(http.MethodGet)
	contactDetailsRouter.HandleFunc("/{id}", controller.UpdateContactDetails).Methods(http.MethodPut)
	contactDetailsRouter.HandleFunc("/{id}", controller.DeleteContactDetails).Methods(http.MethodDelete)
	contactDetailsRouter.Use(auth.IsUser)
	contactDetailsRouter.Use(useridtokenuseridverification.CheckUserVerify)
	fmt.Println("[Contact Details register routes]")

}

// package guru_contact_details

// package guru_contact_details

// import (
// 	contactdetailscontroller "contactsoneapp/components/guru_contact_details/controller"
// 	"contactsoneapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) {
// 	contactDetailsRouter := parentRouter.PathPrefix("/{contact-id}/contact-details").Subrouter()
// 	guidedRouter := contactDetailsRouter.PathPrefix("/g").Subrouter()
// 	guidedRouter.HandleFunc("/", contactdetailscontroller.ReadContactDetailsAll).Methods("GET")
// 	guidedRouter.HandleFunc("/", contactdetailscontroller.CreateContactDetails).Methods("POST")
// 	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.ReadContactDetailsById).Methods("GET")
// 	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.UpdateContactDetailsById).Methods("PUT")
// 	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.DeleteContactDetailsById).Methods("DELETE")
// 	guidedRouter.Use(auth.IsUser)
// }
