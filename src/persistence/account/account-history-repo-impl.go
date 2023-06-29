package account

import (
	"context"
	"database/sql"

	"github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/model"
	"github.com/LintaoAmons/undercontrol/.gen/undercontrol/public/table"
	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	setup "github.com/LintaoAmons/undercontrol/src/persistence/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/pterm/pterm"
)

type AccountHistPostgresRepo struct {
	db     *sql.DB
	logger *pterm.Logger
}

func NewAccountHistoryPostgresRepo() domain.AccountHistoryRepo {
	db := setup.SetupPostgres()
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo).WithMaxWidth(200)
	return &AccountHistPostgresRepo{db: db, logger: logger}
}

func (r *AccountHistPostgresRepo) Insert(tx common.Tx, a *domain.AccountHistory) (string, error) {
	po := AccountConvert(a)

	_, err := table.AccountHistory.
		INSERT(table.AccountHistory.Name,
			table.AccountHistory.Amount,
			table.AccountHistory.CurrencyCode,
			table.AccountHistory.CreatedAt,
			table.AccountHistory.CreatedBy,
			table.AccountHistory.UpdatedAt,
			table.AccountHistory.UpdatedBy).
		MODEL(po).
		ExecContext(context.TODO(), tx.(*sql.Tx))
	if err != nil {
		r.logger.Error(spew.Sdump(err))
		panic(err)
	}

	return po.Name, nil
}

func (r *AccountHistPostgresRepo) FindAllOf(name string) []*domain.AccountHistory {
	x := table.AccountHistory.
		SELECT(table.AccountHistory.AllColumns).WHERE(table.AccountHistory.Name.EQ(postgres.String(name)))

	r.logger.Trace(x.DebugSql())
	var dest []struct {
		model.AccountHistory
	}
	x.Query(r.db, &dest)

	result := []*domain.AccountHistory{}
	for _, v := range dest {
		result = append(result, AccountHistPoToDomain(&v.AccountHistory))
	}

	return result
}
