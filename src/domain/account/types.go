package account

import (
	. "github.com/Rhymond/go-money"
  common "github.com/LintaoAmons/undercontrol/src/domain/common"
)

type Account struct {
	Name   string
	Amount *Money

  common.Audit
}
