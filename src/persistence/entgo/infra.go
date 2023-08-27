package entgo

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/LintaoAmons/undercontrol/ent"
)

var Client *ent.Client

func InitDBClient() *ent.Client {
	if Client == nil {
		c, err := ent.Open(dialect.SQLite, "file:/Users/lintao/Documents/oatnil/release/undercontrol-go/ent-db?mode=memory&mecache=shared&_fk=1")
		// client, err := ent.Open(dialect.SQLite, "file:ent?mecache=shared&_fk=1")
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
