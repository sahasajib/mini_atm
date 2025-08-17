package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func Balance(w http.ResponseWriter, r *http.Request){
	usernameVal := r.Context().Value("username")
	username, ok := usernameVal.(string)
	if !ok {
		http.Error(w, "Unautorized: missing username", http.StatusUnauthorized)
		return
	}
	
	log.Printf("user name %s", username)

	db := database.DB

	var balance float64
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		log.Printf("Database error (fetching user_id): %v", err)
		http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
		return
	}

	// Query total balance from transection table
	
	query := `SELECT total_balance FROM transection WHERE user_id = $1 ORDER BY id DESC LIMIT 1`
	err = db.QueryRow(query, userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			balance = 0.00 // No transactions yet
		} else {
			log.Printf("Database error (fetching total_balance): %v", err)
			http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
			return
		}
	}

	resp := database.BalanceResponse{
		UserName: username,
		Balance:  balance,
	}

	util.SendData(w, resp, http.StatusOK)
}