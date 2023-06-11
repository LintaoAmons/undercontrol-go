package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func SetupPostgres() *sql.DB {
	// TODO: use config as db password and host etc.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	pterm.Info.Printfln(viper.GetString("db.host"))
	pterm.Info.Printfln(viper.GetString("db.password"))
	pterm.Info.Printfln(viper.GetString("db.dbname"))
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	return db
}
