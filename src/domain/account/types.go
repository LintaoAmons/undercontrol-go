package account

import (
	. "github.com/LintaoAmons/go-money"
	common "github.com/LintaoAmons/undercontrol/src/domain/common"
)

type Account struct {
	Name   string
	Amount *Money

	common.Audit
}

type CreateAccountCommand struct {
	Name     string
	Amount   float64
	Currency string
}

type DepositCommand struct {
	Name   string
	Amount float64
}

type (
	WithdarwCommand = DepositCommand
)

type CalibrateCommand struct {
	Name         string
	Amount       float64
	CurrencyCode string
}

type DeleteCommand struct {
	// TODO:
}

type AccountService interface {
	CreateNewAccount(command CreateAccountCommand)
	Calibrate(command CalibrateCommand)
	Withdarw(command WithdarwCommand)
	Deposit(dc DepositCommand)
	Delete(command DeleteCommand)
}

// Read & Write APIs
type AccountRepository interface {
	// TODO: int32 to string. use Name as pk
	Save(tx common.Tx, a *Account) (id int32)

	Get(name string) (*Account, error)
	Find(name string) *Account
	FindAll() []*Account
}

// Only read APIs
type AccountReadRepo interface {
	Get(name string) (*Account, error)
	Find(name string) *Account
	FindAll() []*Account
}

type AccountHistory struct {
	Id int
	Account
}

type AccountHistoryRepo interface {
	Save(tx common.Tx, a *AccountHistory) (string, error)

	FindAllOf(name string) []*AccountHistory
}

func (a *Account) ToHistory(id *int) *AccountHistory {
	var idValue int
	if id == nil {
		idValue = 0
	} else {
		idValue = *id
	}
	return &AccountHistory{
		Id:      idValue,
		Account: *a,
	}
}
