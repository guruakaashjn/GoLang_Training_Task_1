package controller

import (
	"contactsoneapp/components/guru_user/service"
	"contactsoneapp/components/log"
	"contactsoneapp/middleware/auth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	log     log.Log
	service *service.UserService
}

func NewUserController(userService *service.UserService, log log.Log) *UserController {
	return &UserController{
		service: userService,
		log:     log,
	}
}

func (controller *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin-priv", controller.RegisterAdmin).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", controller.RegisterUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/", controller.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", controller.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	userRouter.HandleFunc("/{id}", controller.GetUserById).Methods(http.MethodGet)
	userRouter.Use(auth.IsAdmin)
	fmt.Println("[User register routes]")
}

// package guru_user

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
