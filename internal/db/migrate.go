package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateConn() {
	dbUser :=  "postgres"
	dbPass := "eunice99x"
	dbName := "dblearn"
	dbHost := "localhost"
	dbPort := 5432

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	m, err := migrate.New(
		"file://internal/migration",
		dbURL,
	)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}