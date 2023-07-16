package account_test

import (
	"testing"

	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	. "github.com/LintaoAmons/undercontrol/src/persistence/account"
	"github.com/davecgh/go-spew/spew"
)

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
