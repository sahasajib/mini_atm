package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user database.User
	decoder :=  json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	query :=  `INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id`
	err = database.DB.QueryRow(query, user.Name, string(hash)).Scan(&user.ID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = "" // Clear password before sending response
	util.SendDate(w, user, http.StatusCreated)
}