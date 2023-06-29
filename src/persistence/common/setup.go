package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SetupPostgres() *sql.DB {
	// TODO: use config as db password and host etc.
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		"postgres",
		"password",
		"undercontrol",
		"localhost",
		"5432",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	return db
}
