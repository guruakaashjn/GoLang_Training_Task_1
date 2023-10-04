package controller

import (
	"contactsoneapp/components/guru_user/service"
	"contactsoneapp/components/log"
	"contactsoneapp/errors"
	"contactsoneapp/middleware/auth"
	"contactsoneapp/models/user"
	"contactsoneapp/web"
	"fmt"
	"net/http"
	"strconv"

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

func (controller *UserController) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	newUser := user.User{}
	err := web.UnmarshalJSON(r, &newUser)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	newUser.IsAdmin = true

	err = controller.service.CreateUser(&newUser)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newUser)

}

func (controller *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUser := user.User{}
	err := web.UnmarshalJSON(r, &newUser)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	err = controller.service.CreateUser(&newUser)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newUser)

}

func (controller *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := &[]user.User{}
	var totalCount int
	err := controller.service.GetAllUsers(allUsers, &totalCount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allUsers)
}
func (controller *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	requiredUser := user.User{}
	slugs := mux.Vars(r)

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	err = controller.service.GetUserById(&requiredUser, idTemp)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredUser)
}

func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserToUpdate")
	userToUpdate := user.User{}

	fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &userToUpdate)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)

	intID, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	userToUpdate.ID = uint(intID)
	fmt.Println("User To update")
	fmt.Println(&userToUpdate)
	err = controller.service.UpdateUser(&userToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	web.RespondJSON(w, http.StatusOK, userToUpdate)
}

func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	controller.log.Print("Delete user Call")
	userToDelete := user.User{}
	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	userToDelete.ID = uint(intID)
	err = controller.service.DeleteUser(&userToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete user successfull.")

}

// package controller

// import (
// 	userservice "contactsoneapp/components/guru_user/service"
// 	"contactsoneapp/guru_errors"
// 	"contactsoneapp/middleware/auth"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// // var admin *userservice.User = &userservice.User{

// // 	UserId:    uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
// // 	FirstName: "Admin Initial",
// // 	LastName:  "Admin Surname",
// // 	IsAdmin:   true,
// // 	IsActive:  true,
// // 	// Contacts: append(ContactsTempList, ContactsTempItem),
// // 	Contacts: make([]*contacts_service.Contact, 0),
// // }

// func GetAdminObjectFromCookie(w http.ResponseWriter, r *http.Request) (requiredAdmin *userservice.User) {
// 	defer func() {
// 		if details := recover(); details != nil {
// 			fmt.Println(details)
// 		}
// 	}()
// 	cookie, err1 := r.Cookie("auth")
// 	if err1 != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(err1)
// 		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
// 	}
// 	token := cookie.Value
// 	payload, err2 := auth.Verify(token)
// 	if err2 != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(err2)
// 		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
// 	}
// 	requiredAdminTemp := userservice.ReadUserByUserName(payload.UserName)
// 	if requiredAdminTemp == nil {
// 		json.NewEncoder(w).Encode(guru_errors.AdminDeleted)
// 	}
// 	requiredAdmin = requiredAdminTemp
// 	panic(guru_errors.NewUserError(guru_errors.AdminObjectFound).GetSpecificMessage())

// }

// func CreateAdmin(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Create Admin Controller Function")
// 	var newUserTemp *userservice.User
// 	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewUserError(guru_errors.CreateAdminFailed).GetSpecificMessage())
// 	}
// 	// fmt.Println()
// 	newUser := userservice.CreateAdmin(
// 		newUserTemp.FirstName,
// 		newUserTemp.LastName,
// 		newUserTemp.UserName,
// 		newUserTemp.Password,
// 	)
// 	json.NewEncoder(w).Encode(newUser)
// 	// allUsers := userservice.Users
// 	// json.NewEncoder(w).Encode(allUsers)
// 	panic(guru_errors.NewUserError(guru_errors.CreateAdminSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside Create Admin Controller Function post request done")

// }
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Create User Controller Function")
// 	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
// 	var newUserTemp *userservice.User
// 	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewUserError(guru_errors.CreateUserFailed).GetSpecificMessage())
// 	}
// 	newUser := adminObject.CreateUser(
// 		newUserTemp.FirstName,
// 		newUserTemp.LastName,
// 		newUserTemp.UserName,
// 		newUserTemp.Password,
// 	)
// 	json.NewEncoder(w).Encode(newUser)
// 	panic(guru_errors.NewUserError(guru_errors.CreateUserSuccess).GetSpecificMessage())
// 	// fmt.Println("Leaving Create User controller Function post request done")

// }

// func ReadUserById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read User By Id Controller Function")
// 	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
// 	slugs := mux.Vars(r)
// 	userIdTemp := uuid.MustParse(slugs["user-id"])
// 	var requiredUser *userservice.User = adminObject.ReadUserById(userIdTemp)
// 	if requiredUser == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadUserFailed)
// 		panic(guru_errors.NewUserError(guru_errors.ReadUserFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(requiredUser)
// 	panic(guru_errors.NewUserError(guru_errors.ReadUserSuccess).GetSpecificMessage())
// 	// fmt.Println("Leaving Read User By Id Controller Function get request done")
// }

// func ReadUsersAll(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Read All Users Controller function")
// 	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
// 	var requiredUsersList []*userservice.User = adminObject.ReadAllUsers()
// 	if requiredUsersList == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadUserFailed)
// 		panic(guru_errors.NewUserError(guru_errors.ReadUserFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(requiredUsersList)
// 	panic(guru_errors.NewUserError(guru_errors.ReadUserSuccess).GetSpecificMessage())
// 	// fmt.Println("Leaving Read All Users Controller function get request done")

// }
// func UpdateUserById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Update User By Id Controller function")
// 	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
// 	slugs := mux.Vars(r)
// 	var newUserTemp *userservice.User
// 	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewUserError(guru_errors.UpdateUserFailed).GetSpecificMessage())
// 	}
// 	newUpdatedUser := adminObject.UpdateUserObject(uuid.MustParse(slugs["user-id"]), newUserTemp)
// 	// fmt.Println("Inside UpdatedUserById (updated user): ", newUpdatedUser)
// 	json.NewEncoder(w).Encode(newUpdatedUser)
// 	panic(guru_errors.NewUserError(guru_errors.UpdateUserSuccess).GetSpecificMessage())
// 	// fmt.Println("Leaving Update User By Id Controller function put request done")
// }
// func DeleteUserById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside Delete User By Id Controller function")
// 	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
// 	slugs := mux.Vars(r)
// 	deletedUser := adminObject.DeleteUser(uuid.MustParse(slugs["user-id"]))
// 	json.NewEncoder(w).Encode(deletedUser)
// 	panic(guru_errors.NewUserError(guru_errors.DeleteUserSuccess).GetSpecificMessage())
// 	// fmt.Println("Leaving Delete User By Id Controller function delete request done")
// }
