package handler

import (
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)


func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the cookie by setting MaxAge=-1
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // immediately expire the cookie
	})

	
	res := database.Messages{
		Message: "You have been logged out successfully.",
		Options: []string{"login"}}
	util.SendData(w, res, http.StatusOK)
}