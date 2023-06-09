package account_test

import (
	"testing"

	. "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/Rhymond/go-money"
	"github.com/davecgh/go-spew/spew"
)

func Test_NewAccount(t *testing.T) {
	a := Account{
		Name:   "Test",
		Amount: money.New(1234, money.CNY),
		Audit:  common.DefaultAudit(),
	}

	spew.Dump(a)
}
