package account_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
)

func Test_Decimal(t *testing.T) {
	d := decimal.NewFromFloat32(3.234)
	fmt.Println(d.String())
	dd, _ := decimal.NewFromString(d.String())
	spew.Dump(dd)
}
