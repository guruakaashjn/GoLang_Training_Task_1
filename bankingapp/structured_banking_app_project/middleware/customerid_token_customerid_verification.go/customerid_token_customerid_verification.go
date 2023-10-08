package useridtokenuseridverification

import (
	"bankingapp/middleware/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CheckUserVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("authone")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		token := cookie.Value
		payload, err := auth.Verify(token)
		fmt.Println(payload)
		if err != nil {
			fmt.Println("verify(token)")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		if payload.IsAdmin {
			fmt.Println("payload.IsAdmin")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Admin")
			return

		}

		slugs := mux.Vars(r)
		// fmt.Println(slugs["user-id"])
		// fmt.Println(payload.UserId)
		idTemp, _ := strconv.Atoi(slugs["customer-id"])
		if payload.CustomerId != uint(idTemp) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Error")
			return
		}

		next.ServeHTTP(w, r)
	})
}
