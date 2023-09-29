package guru_user

import (
	// usercontroller "contactsoneapp/components/guru_user/controller"
	usercontroller "contactsoneapp/components/guru_user/controller"

	"github.com/gorilla/mux"
)

func HandleRouter(parentRouter *mux.Router) *mux.Router {
	parentRouter.HandleFunc("/admin", usercontroller.CreateAdmin).Methods("POST")
	userRouter := parentRouter.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/", usercontroller.ReadUsersAll).Methods("GET")
	userRouter.HandleFunc("/", usercontroller.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{user-id}", usercontroller.ReadUserById).Methods("GET")
	userRouter.HandleFunc("/{user-id}", usercontroller.UpdateUserById).Methods("PUT")
	userRouter.HandleFunc("/{user-id}", usercontroller.DeleteUserById).Methods("DELETE")
	return userRouter

}
