package guru_contact_details

import (
	contactdetailscontroller "contactsoneapp/components/guru_contact_details/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) {
	contactDetailsRouter := parentRouter.PathPrefix("/{contact-id}/contact-details").Subrouter()
	contactDetailsRouter.HandleFunc("/", contactdetailscontroller.ReadContactDetailsAll).Methods("GET")
	contactDetailsRouter.HandleFunc("/", contactdetailscontroller.CreateContactDetails).Methods("POST")
	contactDetailsRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.ReadContactDetailsById).Methods("GET")
	contactDetailsRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.UpdateContactDetailsById).Methods("PUT")
	contactDetailsRouter.HandleFunc("/{contact-details-id}", contactdetailscontroller.DeleteContactDetailsById).Methods("DELETE")
}
