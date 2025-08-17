package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sahasajib/mini_atm/database"
	"github.com/sahasajib/mini_atm/util"
)

func Withdraw(w http.ResponseWriter, r *http.Request){
	usernameVal := r.Context().Value("username")
	username, ok := usernameVal.(string)
	if !ok {
		http.Error(w, "Unautorized: missing username", http.StatusUnauthorized)
		return
	}
	log.Printf("user name %s", username)

	var requestMoney database.TransactionRequst

	if err := json.NewDecoder(r.Body).Decode(&requestMoney); err != nil{
		log.Printf("Error decoding json: %v", err)
		http.Error(w, "Invaild request body", http.StatusBadRequest)
		return
	}
	if requestMoney.Amount <= 0{
		http.Error(w, "Amount must be positive", http.StatusBadRequest)
		return
	}
	if requestMoney.Amount > 10000{
		http.Error(w, "Amoun withdral amount 10000", http.StatusBadRequest)
		return
	}
	db := database.DB
	
	tx, err := db.Begin()
	if err != nil{
		log.Printf("Error starting transaction: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var userID int
	err = tx.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil{
		log.Printf("Error fatching userID: %v", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var currentBalance float64
	err = tx.QueryRow("SELECT total_balance FROM transection WHERE user_id = $1 ORDER BY id DESC LIMIT 1", userID).Scan(&currentBalance)
		if err == sql.ErrNoRows {
			currentBalance = 0 // new user with no transactions yet
		} else if err != nil {
			log.Printf("Error fetching balance: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

	if (currentBalance - 500) < requestMoney.Amount {
		http.Error(w, "Insufficent balance", http.StatusBadRequest)
		return
	}
	newBalance := currentBalance - requestMoney.Amount

	// Update users.balance
	_, err = tx.Exec("UPDATE transection SET total_balance = $1 WHERE id = $2", newBalance, userID)
	if err != nil {
		log.Printf("Error updating balance: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO transection (user_id, transactionInfo, balance, total_balance) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err = tx.QueryRow(query, userID, "Withdraw", requestMoney.Amount, newBalance).Scan(&id)
	if err != nil{
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
	response := database.Resposes{Response: "Withdraw successful!"}
	util.SendData(w, response, http.StatusCreated)

}