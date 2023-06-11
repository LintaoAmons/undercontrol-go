package account_test

import (
	"testing"

	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	. "github.com/LintaoAmons/undercontrol/src/persistence/account"
	"github.com/Rhymond/go-money"
	"github.com/davecgh/go-spew/spew"
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

	result, _ := repo.Get("Test3")

	spew.Dump(result)
}

func Test_FindAll(t *testing.T) {
	repo := initRepo()

	result := repo.FindAll()

	spew.Dump(result)
}

func initRepo() domain.AccountRepository {
	return NewAccountRepository()
}
