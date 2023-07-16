package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//go:generate gonstructor --type=PostgresConfigs --constructorTypes=builder
type PostgresConfigs struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
	sslmode  string
}

func SetupPostgres() *sql.DB {
	// TODO: use config as db password and host etc.go install golang.org/x/tools/cmd/goimports@latest
	// viper.GetString("db.password") Why viper can't get the configuration here
	config := NewPostgresConfigsBuilder().
		User("postgres").
		Password("password").
		Dbname("undercontrol").
		Host("localhost").
		Port("5432").Build()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.user,
		config.password,
		config.dbname,
		config.host,
		config.port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	return db
}
