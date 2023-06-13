package moneyutil_test

import (
	"fmt"
	"testing"

	"github.com/LintaoAmons/undercontrol/src/utils/moneyutil"
	money "github.com/Rhymond/go-money"
	"github.com/davecgh/go-spew/spew"
)

func Test_Sum(t *testing.T) {
	m1 := money.New(100, "CNY")
	m2 := money.New(200, "SGD")
	m3 := money.New(300, "USD")
	expectedSum := money.New(3250, "CNY")

	actualSum := moneyutil.Sum("CNY", m1, m2, m3)

	equal, _ := actualSum.Equals(expectedSum)
	if !equal {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", actualSum.Display(), expectedSum.Display())
	}
}

func Test_Convert(t *testing.T) {
	sgdMoney := money.NewFromFloat(100, "SGD")
	cny := moneyutil.NewConvertableMoney(sgdMoney).Convert("CNY")
	equal, _ := cny.Equals(money.NewFromFloat(530, "CNY"))
	if !equal {
		t.Error("Convert error")
	}
}

func Test_MoneyAdd(t *testing.T) {
	m1 := money.NewFromFloat(100.23, "CNY") // 单位元
	m2 := money.New(10023, "CNY")           // 单位分

	m3, _ := m1.Add(m2)

	fmt.Println("m1: " + m1.Display()) // m1: 100.23 元
	fmt.Println("m2: " + m2.Display()) // m2: 100.23 元
	fmt.Println("m3: " + m3.Display()) // m3: 200.46 元
	spew.Dump(m3)
}

func TEST_Money(t *testing.T) {
	m1 := money.New(100, "CNY")
	m1.Multiply(10)
	spew.Dump(m1)
}
