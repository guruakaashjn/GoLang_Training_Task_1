package controller

import (
	"contactsoneapp/guru_errors"
	"contactsoneapp/middleware/auth"
	"contactsoneapp/models/user"
	"contactsoneapp/utils"
	"contactsoneapp/web"
	"encoding/json"
	"net/http"
	"time"
)

func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if details := recover(); details != nil {
			json.NewEncoder(w).Encode(details)
		}
	}()
	var claims *auth.Claims
	err := json.NewDecoder(r.Body).Decode(&claims)
	if err != nil {
		panic(err)
	}
	requiredUser := &user.User{}

	err = controller.service.AuthService(requiredUser, claims.UserName)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	if requiredUser.IsAdmin {
		claims.IsAdmin = true
	}
	claims.UserId = requiredUser.ID

	flag := utils.IsHashMatch([]byte(requiredUser.Password), claims.Password)
	// fmt.Println("Flag: ", flag)
	if !flag {
		panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationFailed).GetSpecificMessage())
	}
	// fmt.Println(guru_errors.NewAuthenticationError(guru_errors.AuthenticationSuccess).GetSpecificMessage())

	token, err := auth.Sign(*claims)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})
	json.NewEncoder(w).Encode("LoggedIn successfully")
	panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationSuccess).GetSpecificMessage())
}
