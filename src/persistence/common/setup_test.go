package common_test

import (
	"testing"

	"github.com/LintaoAmons/undercontrol/src/persistence/common"
)

func Test_SetupPostgres(t *testing.T) {
	db := common.SetupPostgres()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Error("Can't setup postgres")
	}
}
