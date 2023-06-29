package account

import (
	"github.com/LintaoAmons/go-money"
	b "github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
	a "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/shopspring/decimal"
)

func AccountConvert(i *a.Account) *b.Account {
	amount := i.Amount.GetAmount().String()
	currencyCode := i.Amount.Currency().Code
	return &b.Account{
		ID:           0, // ID won't be used when INSERT OR UPDATE
		Name:         i.Name,
		Amount:       &amount,
		CurrencyCode: &currencyCode,

		CreatedAt: &i.CreatedAt,
		CreatedBy: &i.CreatedBy,
		UpdatedAt: &i.UpdatedAt,
		UpdatedBy: &i.UpdatedBy,
	}
}

func AccountPoToDomain(i *b.Account) *a.Account {
	amount, _ := decimal.NewFromString(*i.Amount) // TODO: check error
	money := money.New(amount, *i.CurrencyCode)
	return &a.Account{
		Name:   i.Name,
		Amount: money,
		Audit: common.Audit{
			CreatedAt: *i.CreatedAt,
			CreatedBy: *i.CreatedBy,
			UpdatedAt: *i.UpdatedAt,
			UpdatedBy: *i.UpdatedBy,
		},
	}
}

func AccountHistPoToDomain(i *b.AccountHistory) *a.AccountHistory {
	amount, _ := decimal.NewFromString(*i.Amount) // TODO: check error
	money := money.New(amount, *i.CurrencyCode)
	return &a.Account{
		Name:   i.Name,
		Amount: money,
		Audit: common.Audit{
			CreatedAt: *i.CreatedAt,
			CreatedBy: *i.CreatedBy,
			UpdatedAt: *i.UpdatedAt,
			UpdatedBy: *i.UpdatedBy,
		},
	}
}
