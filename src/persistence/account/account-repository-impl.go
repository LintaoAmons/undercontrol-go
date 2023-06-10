package account

import (
	"database/sql"

	"github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
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
	po := AccountConvert(a)
	x := table.Account.SELECT(table.Account.Name).WHERE(table.Account.Name.EQ(postgres.String(po.Name)))
	ar.logger.Trace(x.DebugSql())
	dest := []string{}
	x.Query(ar.db, &dest)
	if len(dest) > 0 {
		ar.Update(a)
	} else {
		ar.Insert(a)
	}
	return 1
}

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
	_, err := table.Account.
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

	return po.ID, nil
}

func (ar *AccountRepositoryImpl) Get(name string) (*domain.Account, error) {
	x := table.Account.
		SELECT(table.Account.AllColumns).
		WHERE(table.Account.Name.EQ(postgres.String(name)))
	ar.logger.Trace(x.DebugSql())

	var dest []struct {
		model.Account
	}
	x.Query(ar.db, &dest)

	return AccountPoToDomain(&dest[0].Account), nil
}

func (ar *AccountRepositoryImpl) Find(name string) *domain.Account {
	// TODO:
	return nil
}

func (ar *AccountRepositoryImpl) FindAll() *[]domain.Account {
	x := table.Account.
		SELECT(table.Account.AllColumns)

	ar.logger.Trace(x.DebugSql())
	var dest []struct {
		model.Account
	}
	x.Query(ar.db, &dest)

	result := []domain.Account{}
	for _, v := range dest {
		result = append(result, *AccountPoToDomain(&v.Account))
	}

	return &result
}
