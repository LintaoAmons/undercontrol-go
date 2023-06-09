package account

import (
	b "github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
	a "github.com/LintaoAmons/undercontrol/src/domain/account"
)

func AccountConvert(i *a.Account) *b.Account {
	amount := i.Amount.Absolute().Amount()
	currencyCode := i.Amount.Currency().Code
	return &b.Account{
		ID:           1,
		Name:         i.Name,
		Amount:       &amount,
		CurrencyCode: &currencyCode,

		CreatedAt: &i.CreatedAt,
		CreatedBy: &i.CreatedBy,
		UpdatedAt: &i.UpdatedAt,
		UpdatedBy: &i.UpdatedBy,
	}
}
