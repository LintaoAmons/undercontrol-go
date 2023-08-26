package entgo

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/ent"
	"github.com/LintaoAmons/undercontrol/ent/account"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

type AccountEntRepo struct {
	client *ent.Client
}

func NewAccountEntRepo(dsn string) domain.AccountRepository {
	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &AccountEntRepo{
		client: client,
	}
}

func (repo *AccountEntRepo) Save(tx common.Tx, a *domain.Account) (id int32) {
	repo.client.Account.Create().
		SetName(a.Name).SetAmount(a.Amount.GetAmount().String()).
		SetCurrencyCode(a.Amount.Currency().Code).
		SetCreatedAt(a.CreatedAt).
		SetCreatedBy(a.CreatedBy).
		SetUpdatedAt(a.UpdatedAt).
		SetUpdatedBy(a.UpdatedBy).
		SaveX(context.Background())
	// TODO: remove the return value
	return 0
}
func (repo *AccountEntRepo) Get(name string) (*domain.Account, error) {
	ac := repo.client.Account.Query().
		Where(account.NameEQ(name)).
		OnlyX(context.Background())
	return toDomain(ac), nil
}

func (repo *AccountEntRepo) Find(name string) *domain.Account {
	ac := repo.client.Account.Query().
		Where(account.NameEQ(name)).
		OnlyX(context.Background())
	return toDomain(ac)
}

func (repo *AccountEntRepo) FindAll() []*domain.Account {
	acList := repo.client.Account.Query().AllX(context.Background())
	var domains []*domain.Account
	for _, ac := range acList {
		domains = append(domains, toDomain(ac))
	}
	return domains
}

func toDomain(ac *ent.Account) *domain.Account {
	amount, _ := decimal.NewFromString(*ac.Amount) // TODO: check error
	money := money.New(amount, *ac.CurrencyCode)

	return &domain.Account{
		Name:   ac.Name,
		Amount: money,
		Audit: common.Audit{
			CreatedAt: *ac.CreatedAt,
			CreatedBy: *ac.CreatedBy,
			UpdatedAt: *ac.UpdatedAt,
			UpdatedBy: *ac.UpdatedBy,
		},
	}
}
