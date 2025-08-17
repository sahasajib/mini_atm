package database

import (
	"database/sql"
	"log"
	"log/slog"
	_ "github.com/lib/pq"
)

var DB *sql.DB
func InitDB(){
	var err error
	connString := "host=localhost port=5432 user=postgres password=12345678 dbname=mini_atm sslmode=disable"
	DB, err = sql.Open("postgres", connString)
    if err != nil{
		log.Fatal("Failed to connect to the database:", err)
	}
    err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	} 
	slog.Info("Database connection established successfully")

	//user table
	createTableUser :=` CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(100) NOT NULL
	);`
	_, err = DB.Exec(createTableUser)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
	slog.Info("Users table created successfully")

	//trangiction table
	createTableTransection :=` CREATE TABLE IF NOT EXISTS transection (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		transactionInfo VARCHAR(100) DEFAULT 'Main Balance',
		balance DECIMAL(10,2) DEFAULT 0.00,
		total_balance DECIMAL(10,2) DEFAULT 0.00,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = DB.Exec(createTableTransection)
	if err != nil {
		log.Fatal("Failed to create Transection table:", err)
	}
	slog.Info("Transection table created successfully")
	

}