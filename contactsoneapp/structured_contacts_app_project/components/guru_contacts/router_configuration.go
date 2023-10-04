package guru_contacts

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
