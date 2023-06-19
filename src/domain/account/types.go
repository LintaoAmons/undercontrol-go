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

type AccountRepository interface {
	// TODO: int32 to string. use Name as pk
	Insert(a *Account) (int32, error)
	Update(a *Account) (int32, error)
	Save(a *Account) (id int32)

	Get(name string) (*Account, error)
	Find(name string) *Account
	FindAll() []*Account
}
