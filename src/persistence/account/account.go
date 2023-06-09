package account

import (
	"time"

	b "github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
	a "github.com/LintaoAmons/undercontrol/src/domain/account"
)

func AccountConvert(i *a.Account) *b.Account {
	operator := "TODO"
	now := time.Now()
	amount := i.Amount.Absolute().Amount()
	currencyCode := i.Amount.Currency().Code
	return &b.Account{
		ID:           1,
		Name:         i.Name,
		Amount:       &amount,
		CurrencyCode: &currencyCode,

		// TODO: audit fields
		CreatedAt: &now,
		CreatedBy: &operator,
		UpdatedAt: &now,
		UpdatedBy: &operator,
	}
}
