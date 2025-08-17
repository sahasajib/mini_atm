package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)



var jwtKey = []byte("your_secret_key")

func GenerateJWT(userId int, username string)(string, error){
	claims := jwt.MapClaims{
		"user_id" : userId,
		"username": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}