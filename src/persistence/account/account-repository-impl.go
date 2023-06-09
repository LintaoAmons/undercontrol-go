package account

import (
	"database/sql"

	"github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/table"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

type AccountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &AccountRepositoryImpl{db: db}
}

func (ar *AccountRepositoryImpl) Save(a *domain.Account) (id int32) {
	// po := AccountConvert(a)
	// TODO:
	return 1
}

func (ar *AccountRepositoryImpl) Insert(a *domain.Account) (int32, error) {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithMaxWidth(200)

	po := AccountConvert(a)
	result, err := table.Account.
		INSERT(table.Account.Name,
			table.Account.Amount,
			table.Account.CurrencyCode,
			table.Account.CreatedAt,
			table.Account.CreatedBy,
			table.Account.UpdatedAt,
			table.Account.UpdatedBy).
		MODEL(po).
		Exec(ar.db)

	if err != nil {
		logger.Error(spew.Sdump(err))
		panic(err)
	}
	logger.Info(spew.Sdump(result))

	return po.ID, nil
}

func (ar *AccountRepositoryImpl) Get(id string) (*domain.Account, error) {
	// TODO:
	return nil, nil
}

func (ar *AccountRepositoryImpl) Find(id string) *domain.Account {
	// TODO:
	return nil
}

func (ar *AccountRepositoryImpl) FindAll() *[]domain.Account {
	// TODO:
	return nil
}
