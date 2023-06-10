package account_test

import (
	"testing"

	"github.com/LintaoAmons/undercontrol/src/domain/account"
	. "github.com/LintaoAmons/undercontrol/src/persistence/account"
	"github.com/Rhymond/go-money"
	"github.com/davecgh/go-spew/spew"
)

func Test_AccountConvert(t *testing.T) {
	result := AccountConvert(&account.Account{
		Name:   "Lintao",
		Amount: money.New(1234, money.CNY),
	})

	spew.Dump(result)
  spew.Dump(result.Amount)
	spew.Printf("result: %v", result)
	t.Log(result)
}

func Test_Money(t *testing.T) {
  a:=money.New(1234 , money.CNY)
  spew.Dump(a.Absolute())
  spew.Dump(a.Amount())
  spew.Dump(a.Currency())
}

func Test_AccountPoToDomain(t *testing.T) {
  
}
