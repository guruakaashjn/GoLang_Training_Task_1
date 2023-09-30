package controller

import (
	user_service "contactsoneapp/components/guru_user/service"
	"contactsoneapp/guru_errors"
	"contactsoneapp/middleware/auth"
	"contactsoneapp/utils"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	requiredUser := user_service.ReadUserByUserName(claims.UserName)
	// fmt.Println("required user: ", requiredUser)
	if requiredUser == nil {
		panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationFailed).GetSpecificMessage())
	}
	if requiredUser.IsAdmin {
		claims.IsAdmin = true
	}

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
