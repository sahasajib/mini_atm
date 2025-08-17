package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	
)

func Balance(w http.ResponseWriter, r *http.Request){
	username := r.Context().Value("username").(string)

	db := database.DB

	var balance float64
	query := `SELECT balance FROM transactions
	          WHERE username = $1
	          ORDER BY created_at DESC
	          LIMIT 1`
	err := db.QueryRow(query, username).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			balance = 0.00
		} else {
			http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
			return
		}
	}

	resp := map[string]interface{}{
		"username": username,
		"balance":  balance,
	}

	json.NewEncoder(w).Encode(resp)
}