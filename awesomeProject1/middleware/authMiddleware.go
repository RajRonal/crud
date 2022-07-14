package middleware

import (
	"awesomeProject1/handlers"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Header.Get("Authorization")
		token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(handlers.JwtKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			_, err2 := w.Write([]byte("Unauthorized"))
			if err2 != nil {
				return
			}
		}

	})
}
