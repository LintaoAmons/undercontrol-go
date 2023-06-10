package account

import (
	b "github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
	a "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/Rhymond/go-money"
)

func AccountConvert(i *a.Account) *b.Account {
	amount := i.Amount.Absolute().Amount()
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
	money := money.New(*i.Amount, *i.CurrencyCode)
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
