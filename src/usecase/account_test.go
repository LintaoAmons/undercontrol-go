package usecase_test

import (
	"testing"

	"github.com/LintaoAmons/go-money"
	"github.com/davecgh/go-spew/spew"
)

func Test_MoneyFloat(t *testing.T) {
	a := money.NewFromFloat(1234.12341234, "CNY")
	spew.Dump(a)
}
