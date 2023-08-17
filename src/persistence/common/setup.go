package common

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type PostgresConfigs struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

func SetupPostgres() *sql.DB {
	// TODO: use config as db password and host etc.go install golang.org/x/tools/cmd/goimports@latest
	// viper.GetString("db.password") Why viper can't get the configuration here
	var config PostgresConfigs
	config.GetConf()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User,
		config.Password,
		config.Dbname,
		config.Host,
		config.Port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	return db
}

func (c *PostgresConfigs) GetConf() *PostgresConfigs {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	// TODO: filepath as env
	filePath := usr.HomeDir + "/.undercontrol.yaml"
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
