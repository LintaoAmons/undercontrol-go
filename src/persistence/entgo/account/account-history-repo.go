package account

import (
	"context"

	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/ent"
	"github.com/LintaoAmons/undercontrol/ent/accounthistory"
	"github.com/LintaoAmons/undercontrol/src/domain/account"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/shopspring/decimal"
)

type AccountHistoryEntRepo struct {
	client *ent.Client
}

func NewAccountHistoryEntRepo(c *ent.Client) domain.AccountHistoryRepo {
	return &AccountHistoryEntRepo{
		client: c,
	}
}

func (r *AccountHistoryEntRepo) Save(tx common.Tx, a *account.AccountHistory) (string, error) {
	r.client.AccountHistory.Create().
		SetName(a.Name).SetAmount(a.Amount.GetAmount().String()).
		SetCurrencyCode(a.Amount.Currency().Code).
		SetCreatedAt(a.CreatedAt).
		SetCreatedBy(a.CreatedBy).
		SetUpdatedAt(a.UpdatedAt).
		SetUpdatedBy(a.UpdatedBy).
		SaveX(context.Background())
	// TODO: remove the return value
	return a.Name, nil
}
func (r *AccountHistoryEntRepo) FindAllOf(name string) []*account.AccountHistory {
	acList := r.client.AccountHistory.Query().
		Where(accounthistory.NameEQ(name)).
		AllX(context.Background())
	return toHistDomain(acList)
}

func toHistDomain(histories ent.AccountHistories) []*domain.AccountHistory {
	var domains []*domain.AccountHistory
	for _, h := range histories {
		amount, _ := decimal.NewFromString(*h.Amount)
		money := money.New(amount, *h.CurrencyCode)

		domains = append(domains, &domain.AccountHistory{
			Id: h.ID,
			Account: account.Account{
				Name:   h.Name,
				Amount: money,
				Audit: common.Audit{
					CreatedAt: *h.CreatedAt,
					CreatedBy: *h.CreatedBy,
					UpdatedAt: *h.UpdatedAt,
					UpdatedBy: *h.UpdatedBy,
				},
			},
		},
		)
	}
	return domains
}
