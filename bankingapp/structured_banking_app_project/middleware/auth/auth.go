package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	CustomerId uint
	UserName   string
	Password   string
	IsAdmin    bool
	jwt.StandardClaims
}

var secretKeyForJWT = []byte("bAnKInG1App2345Guru9876")

func Sign(claims Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SignedToken, err := token.SignedString(secretKeyForJWT)
	if err != nil {
		return "", err
	}
	return SignedToken, nil
}

func Verify(token string) (*Claims, error) {
	var claims = &Claims{}
	fmt.Println(token)

	tokenObj, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKeyForJWT, nil
	})
	if err != nil {
		return nil, err
	}
	if !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err1 := r.Cookie("authone")
		if err1 != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err1)
			return
		}
		token := cookie.Value
		payload, err2 := Verify(token)
		if err2 != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err2)
			return
		}
		if !payload.IsAdmin {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not Admin")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err1 := r.Cookie("authone")
		if err1 != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err1)
			return
		}
		token := cookie.Value
		payload, err2 := Verify(token)
		if err2 != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err2)
			return
		}
		if payload.IsAdmin {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Admin")
			return
		}
		next.ServeHTTP(w, r)

	})
}
