package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
	usernameVal := r.Context().Value("username")
	username, ok := usernameVal.(string)
	if !ok {
		http.Error(w, "Unautorized: missing username", http.StatusUnauthorized)
		return
	}

	log.Printf("user name %s", username)

	var requestMoney database.TransactionRequst

	if err := json.NewDecoder(r.Body).Decode(&requestMoney); err != nil {
		log.Printf("Error decoding Json: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadGateway)
		return
	}
	if requestMoney.Amount <= 0 {
		http.Error(w, "Amount must be positive", http.StatusBadRequest)
		return
	}

	if requestMoney.Amount > 10000 {
		http.Error(w, "Amoun deposit amount 10000", http.StatusBadRequest)
		return
	}

	db := database.DB
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var userID int
	err = tx.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		log.Printf("Error fatching userID: %v", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var lastBalance float64
	err = tx.QueryRow("SELECT COALESCE(total_balance, 0) FROM transection WHERE user_id = $1 ORDER BY id DESC LIMIT 1", userID).Scan(&lastBalance)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	newBalance := lastBalance + requestMoney.Amount

	query := `INSERT INTO transection (user_id, transactionInfo, balance, total_balance) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err = tx.QueryRow(query, userID, "Deposit", requestMoney.Amount, newBalance).Scan(&id)
	if err != nil {
		log.Printf("Error inserting transaction: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Send response
	response := database.Resposes{Response: "Deposit successful!"}
	util.SendData(w, response, http.StatusCreated)
}
