package account_test

import (
	"testing"

	"github.com/LintaoAmons/go-money"
	. "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
)

func Test_NewAccount(t *testing.T) {
	a := Account{
		Name:   "Test",
		Amount: money.New(decimal.NewFromInt(1234), money.CNY),
		Audit:  common.DefaultAudit(),
	}

	spew.Dump(a)
}
