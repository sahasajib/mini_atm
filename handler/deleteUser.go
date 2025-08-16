package handler

import (
	"database/sql"
	"net/http"
	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func DelteUser(w http.ResponseWriter, r *http.Request) {
	id, err := util.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Please provide a vaild user Id", http.StatusBadRequest)
	}


	

	db := database.DB

	// Check if user exists
	var user database.User
	err = db.QueryRow("SELECT id, name, password FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to check user", http.StatusInternalServerError)
		return
	}

	// Delete user
	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	util.SendDate(w, user, http.StatusOK)
}
