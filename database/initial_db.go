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


	createTableUser :=` CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		password VARCHAR(100) NOT NULL
	);`
	_, err = DB.Exec(createTableUser)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
	slog.Info("Users table created successfully")
	

}