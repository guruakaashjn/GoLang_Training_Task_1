package guru_user

// package guru_user

// import (
// 	// usercontroller "contactsoneapp/components/guru_user/controller"
// 	usercontroller "contactsoneapp/components/guru_user/controller"
// 	"contactsoneapp/middleware/auth"

// 	"github.com/gorilla/mux"
// )

// func HandleRouter(parentRouter *mux.Router) *mux.Router {
// 	parentRouter.HandleFunc("/admin", usercontroller.CreateAdmin).Methods("POST")
// 	parentRouter.HandleFunc("/login", usercontroller.Login).Methods("POST")

// 	userRouter := parentRouter.PathPrefix("/user").Subrouter()
// 	guardedRouter := userRouter.PathPrefix("/g").Subrouter()
// 	guardedRouter.HandleFunc("/", usercontroller.ReadUsersAll).Methods("GET")
// 	guardedRouter.HandleFunc("/", usercontroller.CreateUser).Methods("POST")
// 	guardedRouter.HandleFunc("/{user-id}", usercontroller.ReadUserById).Methods("GET")
// 	guardedRouter.HandleFunc("/{user-id}", usercontroller.UpdateUserById).Methods("PUT")
// 	guardedRouter.HandleFunc("/{user-id}", usercontroller.DeleteUserById).Methods("DELETE")
// 	guardedRouter.Use(auth.IsAdmin)
// 	return userRouter

// }
