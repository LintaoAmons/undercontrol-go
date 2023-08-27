package account_test

import (
	"fmt"
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

func Test_AccountToHistory(t *testing.T) {
	a := Account{
		Name:   "Test",
		Amount: money.New(decimal.NewFromInt(1234), money.CNY),
		Audit:  common.DefaultAudit(),
	}

	r := a.ToHistory(nil)

	fmt.Println(r.Name)
	fmt.Println(r.Amount)
	fmt.Println(r.Audit)
}
