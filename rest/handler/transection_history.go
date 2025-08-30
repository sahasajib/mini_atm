package handler

import (
	"log"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)


func TransectionHistory(w http.ResponseWriter, r *http.Request){
	usernameVal := r.Context().Value("username")
	username, ok := usernameVal.(string)
	if !ok {
		http.Error(w, "Unauthorized: missing username", http.StatusUnauthorized)
		return
	}
	log.Printf("user name: %s", username)

	db := database.DB

	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		log.Printf("Error fetching userID: %v", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	rows, err := db.Query("SELECT transactioninfo, balance, created_at FROM transection WHERE user_id=$1 ORDER BY created_at DESC LIMIT 10", userID)
	if err != nil {
		log.Printf("Error fetching transaction history: %v", err)
		http.Error(w, "Failed to fetch history", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var history []database.TransectionHistory
	for rows.Next() {
		var t database.TransectionHistory
		if err := rows.Scan(&t.TransectionInfo, &t.Amount, &t.CreatedAt); err != nil {
			log.Printf("Error scanning transaction: %v", err)
			continue
		}
		history = append(history, t)
	}
	if len(history) == 0{
		http.Error(w, "No transection found", http.StatusNotFound)
		return
	}

	util.SendData(w, history, http.StatusOK)
}
