package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("jwt_token")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenStr := cookie.Value
        claims := jwt.MapClaims{}
        token, err :=jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("your_secret_key"), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // context এ username set করা
		username, ok := claims["username"].(string)
		if !ok {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
        ctx := context.WithValue(r.Context(), "username", username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
