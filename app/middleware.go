package middleware

import (
	"context"
	"net/http"

	"github.com/ironline/iapp/api/config"

	"github.com/golang-jwt/jwt"
)

// Authmiddleware for handling JWT authentication
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get token from header
		tokenString := r.Header.Get("Authorization")

		// parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// set user id from token as context
		ctx := context.WithValue(r.Context(), "user_id", token.Claims.(jwt.MapClaims)["sub"])

		// pass context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
