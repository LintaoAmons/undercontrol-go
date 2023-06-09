package account

import (
	common "github.com/LintaoAmons/undercontrol/src/domain/common"
	. "github.com/Rhymond/go-money"
)

type Account struct {
	Name   string
	Amount *Money

	common.Audit
}

type AccountRepository interface {
  Insert(a *Account) (int32, error)
	Save(a *Account) (id int32)

	Get(id string) (*Account, error)
  Find(id string) (*Account)
  FindAll() (*[]Account)
}
