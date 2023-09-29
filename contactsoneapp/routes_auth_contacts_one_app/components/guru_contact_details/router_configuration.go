package guru_contact_details

import (
	contactdetailscontroller "contactsoneapp/components/guru_contact_details/controller"
	"contactsoneapp/middleware/auth"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	contactDetailsRouter := parentRouter.PathPrefix("/{contact-id}/contact-details").Subrouter()
	guidedRouter := contactDetailsRouter.PathPrefix("/g").Subrouter()
	guidedRouter.HandleFunc("/", contactdetailscontroller.ReadContactDetailsAll).Methods("GET")
	guidedRouter.HandleFunc("/", contactdetailscontroller.CreateContactDetails).Methods("POST")
	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.ReadContactDetailsById).Methods("GET")
	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.UpdateContactDetailsById).Methods("PUT")
	guidedRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.DeleteContactDetailsById).Methods("DELETE")
	guidedRouter.Use(auth.IsUser)
}
