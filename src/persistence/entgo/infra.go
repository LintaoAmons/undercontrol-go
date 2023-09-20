package entgo

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/LintaoAmons/undercontrol/ent"
	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func InitDBClient() *ent.Client {
	if Client == nil {
		dbPath := os.Getenv("DATABASE_PATH")
		c, err := ent.Open(dialect.SQLite, "file:"+dbPath+"?mecache=shared&_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		if err := c.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		Client = c
	}

	return Client
}
