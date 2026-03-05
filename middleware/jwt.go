package middleware

import (
	"net/http"
	"strings"
	"user-auth-api/utils"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var tokenString string

		cookie, err := r.Cookie("token")
		if err == nil {
			tokenString = cookie.Value
		} else {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}