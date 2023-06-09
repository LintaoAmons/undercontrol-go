package account_test

import (
	"testing"

	"github.com/LintaoAmons/undercontrol/src/domain/common"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	. "github.com/LintaoAmons/undercontrol/src/persistence/account"
	setup "github.com/LintaoAmons/undercontrol/src/persistence/common"
	"github.com/Rhymond/go-money"
)

func Test_Insert(t *testing.T) {
	repo := NewAccountRepository(setup.SetupPostgres())
	repo.Insert(&domain.Account{
		Name:   "Test3",
		Amount: money.New(1234, money.CNY),
		Audit:  common.DefaultAudit(),
	})
}
