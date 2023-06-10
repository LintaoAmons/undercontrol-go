package account

import (
	"database/sql"

	"github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/table"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/pterm/pterm"
)

type AccountRepositoryImpl struct {
	db     *sql.DB
	logger *pterm.Logger
}

func NewAccountRepository(db *sql.DB, logger *pterm.Logger) domain.AccountRepository {
	return &AccountRepositoryImpl{db: db, logger: logger}
}

func (ar *AccountRepositoryImpl) Save(a *domain.Account) (id int32) {
	// po := AccountConvert(a)
	// TODO:
	return 1
}

// func (p *PocketRepositoryImpl) Update(pocket Pocket) (string, error) {
// 	model := pocketToModel(pocket)
// 	x := table.Pocket.UPDATE(table.Pocket.AllColumns).MODEL(model).
// 		WHERE(table.Pocket.ID.EQ(String(pocket.Id)))
// 	logger.Trace(x.DebugSql())
// 	result, err := x.Exec(p.db)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logger.Trace(fmt.Sprintf("LastInsertId: %v", result))
// 	return pocket.Id, nil
// }

func (ar *AccountRepositoryImpl) Update(a *domain.Account) (int32, error) {
	po := AccountConvert(a)
	x := table.Account.
		UPDATE(
			table.Account.Name,
			table.Account.Amount,
			table.Account.CurrencyCode,
			table.Account.UpdatedAt,
			table.Account.UpdatedBy).
		MODEL(po).
		WHERE(table.Account.Name.EQ(postgres.String(po.Name)))
	ar.logger.Trace(x.DebugSql())
	_, err := x.Exec(ar.db)
	if err != nil {
		ar.logger.Error(spew.Sdump(err))
		return 0, err
	}
	return po.ID, nil
}

func (ar *AccountRepositoryImpl) Insert(a *domain.Account) (int32, error) {
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
		ar.logger.Error(spew.Sdump(err))
		panic(err)
	}
	ar.logger.Info(spew.Sdump(result))

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
