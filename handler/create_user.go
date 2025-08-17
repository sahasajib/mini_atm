package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user database.User
	decoder :=  json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	HashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Filed hash password", http.StatusInternalServerError)
		return
	}

	query :=  `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err = database.DB.QueryRow(query, user.UserName, string(HashPassword)).Scan(&user.ID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = "" // Clear password before sending response
	util.SendData(w, user, http.StatusCreated)
}