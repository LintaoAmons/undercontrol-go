package account_test

import (
	. "github.com/LintaoAmons/undercontrol/src/domain/account"
	common "github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/Rhymond/go-money"
)

func AccountFixture() Account {
  // TODO: How to use fixture in golang
  // https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
	return Account{
		Name:   "Test",
		Amount: money.New(1234, money.CNY),
		Audit:  common.DefaultAudit(),
	}
}
