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
	query := "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id"
	err = database.InitDB().QueryRow(query, user.Name, user.Password).Scan(&user.ID)
	if err != nil{
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	util.SendDate(w, user, http.StatusCreated)
}