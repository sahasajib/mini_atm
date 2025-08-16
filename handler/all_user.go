package handler

import (
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)


func AllUser(w http.ResponseWriter, r *http.Request){
	var users []database.User
	rows, err := database.DB.Query("SELECT id, name, password FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next(){
		var user database.User
		err := rows.Scan(&user.ID, &user.UserName, &user.Password)
		if err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	if len(users) == 0{
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}
	util.SendDate(w, users, http.StatusOK)
}