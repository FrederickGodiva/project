package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=password dbname=ecommerce sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}

func GetDB() *sql.DB {
	return db
}
