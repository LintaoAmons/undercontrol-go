package common

import (
	"database/sql"

	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/pterm/pterm"
)

// TxManagerPostgres
type TxManagerPostgres struct {
	db     *sql.DB
	logger *pterm.Logger
}

func NewTxManagerPostgres() *TxManagerPostgres {
	db := SetupPostgres()
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo).WithMaxWidth(200)
	return &TxManagerPostgres{
		db:     db,
		logger: logger,
	}
}

func (t *TxManagerPostgres) Begin() common.Tx {
	tx, _ := t.db.Begin()
	return tx
}
