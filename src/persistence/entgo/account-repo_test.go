package entgo_test

import (
	"testing"

	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/LintaoAmons/undercontrol/src/persistence/entgo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAccountEntRepo_Save(t *testing.T) {
	repo := entgo.NewAccountEntRepo("file:ent?mode=memory&mecache=shared&_fk=1")
	money := money.New(decimal.NewFromInt(1000), "CNY")

	account := &account.Account{
		Name:   "Test",
		Amount: money,
		Audit:  common.DefaultAudit(),
	}

	// Act
	repo.Save(nil, account)

	// Assert
	savedAccount, err := repo.Get("Test")
	if err != nil {
		t.Fatalf("failed getting account: %v", err)
	}

	assert.Equal(t, account.Name, savedAccount.Name, "Expected Name to be equal")
	assert.Equal(t, account.Amount, savedAccount.Amount, "Expected Amount to be equal")
	assert.Equal(t, account.Audit, savedAccount.Audit, "Expected Audit to be equal")
}
