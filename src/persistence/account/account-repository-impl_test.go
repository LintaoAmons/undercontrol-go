package account_test

import (
	"testing"

	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	. "github.com/LintaoAmons/undercontrol/src/persistence/account"
	setup "github.com/LintaoAmons/undercontrol/src/persistence/common"
	"github.com/Rhymond/go-money"
	"github.com/pterm/pterm"
)

func Test_Save(t *testing.T) {
	repo := initRepo()

	repo.Save(&domain.Account{
		Name:   "Test4",
		Amount: money.New(1234, money.CNY),
		Audit:  common.DefaultAudit(),
	})
}

func Test_Insert(t *testing.T) {
	repo := initRepo()

	repo.Insert(&domain.Account{
		Name:   "Test3",
		Amount: money.New(1234, money.CNY),
		Audit:  common.DefaultAudit(),
	})
}

func Test_Update(t *testing.T) {
	repo := initRepo()

	repo.Update(&domain.Account{
		Name:   "Test3",
		Amount: money.New(343, money.CNY),
		Audit:  common.DefaultAudit(),
	})
}

func Test_Get(t *testing.T) {
	repo := initRepo()

	repo.Get("Test3")
}

func initRepo() domain.AccountRepository {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithMaxWidth(200)

	return NewAccountRepository(setup.SetupPostgres(), logger)
}
