package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func ConnectDB() error {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=fiberlearn sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %v", err)
	}

	DBConnection = db
	fmt.Println("Database connected successfully")
	return nil
}