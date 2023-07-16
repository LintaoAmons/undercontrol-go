package common_test

import (
	"testing"

	"github.com/LintaoAmons/undercontrol/src/persistence/common"
	"github.com/davecgh/go-spew/spew"
)

func Test_SetupPostgres(t *testing.T) {
	db := common.SetupPostgres()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Error("Can't setup postgres")
	}
}

func Test_GetConfig(t *testing.T) {
	var c common.PostgresConfigs
	c.GetConf()
	spew.Dump(c)
}
