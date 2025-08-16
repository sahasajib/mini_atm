package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func GetUser(w http.ResponseWriter, r *http.Request){
	id, err := util.ExtractIDFromPath(r)
	if err != nil{
		http.Error(w, "Please provide a valid user ID", http.StatusBadRequest)
		return
	}
	// database connection
	db := database.DB

	var user database.User
	
	query := "SELECT id, username, password FROM users WHERE id = $1"
	slog.Info("Executing query", "query", query, "id", id)

	err = db.QueryRow(query, id).Scan(&user.ID, &user.UserName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}
	
	util.SendDate(w, user, http.StatusOK)

}