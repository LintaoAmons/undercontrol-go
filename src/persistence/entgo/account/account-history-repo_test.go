package account_test

import (
	"testing"

	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	p "github.com/LintaoAmons/undercontrol/src/persistence/entgo/account"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestAccountEntHistRepo_Save(t *testing.T) {
	repo := p.NewAccountHistoryEntRepo(initDBClient())
	money := money.New(decimal.NewFromInt(1000), "CNY")

	account1 := &account.AccountHistory{
		Id: 0,
		Account: account.Account{
			Name:   "Test",
			Amount: money,
			Audit:  common.DefaultAudit(),
		},
	}

	account2 := &account.AccountHistory{
		Id: 0,
		Account: account.Account{
			Name:   "Test",
			Amount: money,
			Audit:  common.DefaultAudit(),
		},
	}

	// Act
	repo.Save(nil, account1)
	repo.Save(nil, account2)

	// TODO: test

	// Assert
	savedAccounts := repo.FindAllOf("Test")

	assert.Equal(t, len(savedAccounts), 2)
	// assert.Equal(t, account.Name, savedAccount.Name, "Expected Name to be equal")
	// assert.Equal(t, account.Amount, savedAccount.Amount, "Expected Amount to be equal")
	// assert.Equal(t, account.Audit, savedAccount.Audit, "Expected Audit to be equal")
}
