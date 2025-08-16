package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request){
	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("Received user: %+v", user)
	if user.UserName == "" || user.Password == ""{
		http.Error(w, "Username & Password required", http.StatusBadRequest)
		return
	}


	db := database.DB
	var dbUser database.User
	query := "SELECT id, username, password from users WHERE username=$1"
	err := db.QueryRow(query, user.UserName).Scan(&dbUser.ID, &dbUser.UserName, &dbUser.Password)
	if err != nil{
		if err == sql.ErrNoRows{
			http.Error(w, "Invalid username", http.StatusUnauthorized)
			return
		}
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}


	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil{
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	response := database.Messages{Message: "Welcome, " + dbUser.UserName + "! 🎉"}
	util.SendDate(w, response, http.StatusAccepted)

}