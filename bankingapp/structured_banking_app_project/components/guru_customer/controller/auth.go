package controller

import (
	"bankingapp/guru_errors"
	"bankingapp/middleware/auth"
	"bankingapp/models/customer"
	"bankingapp/utils"
	"bankingapp/web"
	"encoding/json"
	"net/http"
	"time"
)

func (controller *CustomerController) Login(w http.ResponseWriter, r *http.Request) {
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
	requiredCustomer := &customer.Customer{}

	err = controller.service.AuthService(requiredCustomer, claims.UserName)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	if requiredCustomer.IsAdmin {
		claims.IsAdmin = true
	}
	claims.CustomerId = requiredCustomer.ID

	flag := utils.IsHashSame(requiredCustomer.Password, claims.Password)
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
		Name:    "authone",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})
	json.NewEncoder(w).Encode("LoggedIn successfully")
	panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationSuccess).GetSpecificMessage())
}
