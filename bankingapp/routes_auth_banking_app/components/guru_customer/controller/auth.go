package controller

import (
	customer_service "bankingapp/components/guru_customer/service"
	"bankingapp/guru_errors"
	"bankingapp/middleware/auth"
	"bankingapp/utils"
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
	claims.IsAdmin = false
	requiredCustomer := customer_service.ReadCustomerByUserName(claims.UserName)

	if requiredCustomer == nil {
		panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationFailed).GetSpecificMessage())
	}

	if requiredCustomer.IsAdmin {
		claims.IsAdmin = true
	}

	flag := utils.IsHashSame(requiredCustomer.Password, claims.Password)

	if !flag {
		panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationFailed).GetSpecificMessage())

	}
	token, err := auth.Sign(*claims)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "authone",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})
	json.NewEncoder(w).Encode("LoggedIn Successfully")
	panic(guru_errors.NewAuthenticationError(guru_errors.AuthenticationSuccess).GetSpecificMessage())

}
