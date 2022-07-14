package main

import (
	"awesomeProject1/Claims"
	"awesomeProject1/Credentials"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var JwtKey = []byte("secret_key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	expirationTime := time.Now().Add(time.Hour * 7)
	claims := &Claims.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	jwt.NewWithClaims(jwt.SigningMethodES256, claims)

}
