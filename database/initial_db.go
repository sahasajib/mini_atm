package database

import (
	"database/sql"
	"log"
	"log/slog"
	_ "github.com/lib/pq"
)

var db *sql.DB
func InitDB()*sql.DB{
	var err error
	connString := "host=localhost port=5432 user=postgres password=12345678 dbname=mini_atm sslmode=disable"
	db, err = sql.Open("postgres", connString)
    if err != nil{
		log.Fatal("Failed to connect to the database:", err)
	}
    err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	} 
	slog.Info("Database connection established successfully")


	createTableUser :=` CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL
	);`
	_, err = db.Exec(createTableUser)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	} else {
		slog.Info("Users table created successfully")
	}
	return db
}