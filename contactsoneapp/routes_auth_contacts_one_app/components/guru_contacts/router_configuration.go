package guru_contacts

import (
	devicecontroller "contactsoneapp/components/guru_contacts/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) *mux.Router {
	contactsRouter := parentRouter.PathPrefix("/{user-id}/contacts").Subrouter()

	contactsRouter.HandleFunc("/", devicecontroller.ReadContactsAll).Methods("GET")
	contactsRouter.HandleFunc("/", devicecontroller.CreateContact).Methods("POST")
	contactsRouter.HandleFunc("/{contact-id}", devicecontroller.ReadContactById).Methods("GET")
	contactsRouter.HandleFunc("/{contact-id}", devicecontroller.UpdateContactById).Methods("PUT")
	contactsRouter.HandleFunc("/{contact-id}", devicecontroller.DeleteContactById).Methods("DELETE")
	return contactsRouter

}
