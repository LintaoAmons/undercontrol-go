package entgo_test

import (
	"context"
	"testing"

	"github.com/LintaoAmons/undercontrol/src/persistence/entgo"
	"github.com/davecgh/go-spew/spew"
)

func Test_InitDbClient(t *testing.T) {
	client := entgo.InitDBClient()

	r := client.Account.Query().AllX(context.Background())

	spew.Dump(r)
}
