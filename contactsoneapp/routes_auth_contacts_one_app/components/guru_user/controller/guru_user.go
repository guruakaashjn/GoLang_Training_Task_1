package controller

import (
	userservice "contactsoneapp/components/guru_user/service"
	"contactsoneapp/guru_errors"
	"contactsoneapp/middleware/auth"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// var admin *userservice.User = &userservice.User{

// 	UserId:    uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
// 	FirstName: "Admin Initial",
// 	LastName:  "Admin Surname",
// 	IsAdmin:   true,
// 	IsActive:  true,
// 	// Contacts: append(ContactsTempList, ContactsTempItem),
// 	Contacts: make([]*contacts_service.Contact, 0),
// }

func GetAdminObjectFromCookie(w http.ResponseWriter, r *http.Request) (requiredAdmin *userservice.User) {
	defer func() {
		if details := recover(); details != nil {
			fmt.Println(details)
		}
	}()
	cookie, err1 := r.Cookie("auth")
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err1)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}
	token := cookie.Value
	payload, err2 := auth.Verify(token)
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err2)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}
	requiredAdminTemp := userservice.ReadUserByUserName(payload.UserName)
	if requiredAdminTemp == nil {
		json.NewEncoder(w).Encode(guru_errors.AdminDeleted)
	}
	requiredAdmin = requiredAdminTemp
	panic(guru_errors.NewUserError(guru_errors.AdminObjectFound).GetSpecificMessage())

}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Create Admin Controller Function")
	var newUserTemp *userservice.User
	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.CreateAdminFailed).GetSpecificMessage())
	}
	// fmt.Println()
	newUser := userservice.CreateAdmin(
		newUserTemp.FirstName,
		newUserTemp.LastName,
		newUserTemp.UserName,
		newUserTemp.Password,
	)
	json.NewEncoder(w).Encode(newUser)
	// allUsers := userservice.Users
	// json.NewEncoder(w).Encode(allUsers)
	panic(guru_errors.NewUserError(guru_errors.CreateAdminSuccess).GetSpecificMessage())
	// fmt.Println("Inside Create Admin Controller Function post request done")

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Create User Controller Function")
	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
	var newUserTemp *userservice.User
	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.CreateUserFailed).GetSpecificMessage())
	}
	newUser := adminObject.CreateUser(
		newUserTemp.FirstName,
		newUserTemp.LastName,
		newUserTemp.UserName,
		newUserTemp.Password,
	)
	json.NewEncoder(w).Encode(newUser)
	panic(guru_errors.NewUserError(guru_errors.CreateUserSuccess).GetSpecificMessage())
	// fmt.Println("Leaving Create User controller Function post request done")

}

func ReadUserById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Read User By Id Controller Function")
	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	userIdTemp := uuid.MustParse(slugs["user-id"])
	var requiredUser *userservice.User = adminObject.ReadUserById(userIdTemp)
	if requiredUser == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadUserFailed)
		panic(guru_errors.NewUserError(guru_errors.ReadUserFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredUser)
	panic(guru_errors.NewUserError(guru_errors.ReadUserSuccess).GetSpecificMessage())
	// fmt.Println("Leaving Read User By Id Controller Function get request done")
}

func ReadUsersAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Read All Users Controller function")
	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
	var requiredUsersList []*userservice.User = adminObject.ReadAllUsers()
	if requiredUsersList == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadUserFailed)
		panic(guru_errors.NewUserError(guru_errors.ReadUserFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredUsersList)
	panic(guru_errors.NewUserError(guru_errors.ReadUserSuccess).GetSpecificMessage())
	// fmt.Println("Leaving Read All Users Controller function get request done")

}
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Update User By Id Controller function")
	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	var newUserTemp *userservice.User
	err := json.NewDecoder(r.Body).Decode(&newUserTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewUserError(guru_errors.UpdateUserFailed).GetSpecificMessage())
	}
	newUpdatedUser := adminObject.UpdateUserObject(uuid.MustParse(slugs["user-id"]), newUserTemp)
	// fmt.Println("Inside UpdatedUserById (updated user): ", newUpdatedUser)
	json.NewEncoder(w).Encode(newUpdatedUser)
	panic(guru_errors.NewUserError(guru_errors.UpdateUserSuccess).GetSpecificMessage())
	// fmt.Println("Leaving Update User By Id Controller function put request done")
}
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside Delete User By Id Controller function")
	var adminObject *userservice.User = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	deletedUser := adminObject.DeleteUser(uuid.MustParse(slugs["user-id"]))
	json.NewEncoder(w).Encode(deletedUser)
	panic(guru_errors.NewUserError(guru_errors.DeleteUserSuccess).GetSpecificMessage())
	// fmt.Println("Leaving Delete User By Id Controller function delete request done")
}
