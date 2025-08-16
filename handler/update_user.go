package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := util.ExtractIDFromPath(r)
	if err != nil || id <= 0 {
		http.Error(w, "Please provide vaild Id for updateding", http.StatusBadRequest)
	}

	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invaild request body", http.StatusBadRequest)
		return
	}

	HashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Filed hash password", http.StatusInternalServerError)
		return
	}

	db := database.DB
	// Check if user exists
	query := "UPDATE users SET name=$1, password=$2 WHERE id=$3"
	res, err := db.Exec(query, user.Name, HashPassword, id)
	if err != nil {
		http.Error(w, "Faild to update data", http.StatusInternalServerError)
		return

	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// Delete user
	user.ID = id
	util.SendDate(w, user, http.StatusOK)
}
