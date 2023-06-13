package moneyutil

import (
	money "github.com/Rhymond/go-money"
	"github.com/shopspring/decimal"
)

type ConvertableMoney struct {
	*money.Money
}

func NewConvertableMoney(m *money.Money) *ConvertableMoney {
	return &ConvertableMoney{
		Money: m,
	}
}

func (m *ConvertableMoney) Convert(currencyCode string) *money.Money {
	exchangeRate := getExchangeRate(m.Currency().Code, currencyCode)
	currentAmount := decimal.NewFromInt(m.Amount()).Div(decimal.NewFromInt(100))
	return money.New(int64(currentAmount.Mul(exchangeRate).Round(2).Mul(decimal.NewFromInt(100)).IntPart()), currencyCode)
}

func Sum(currencyCode string, m ...*money.Money) *money.Money {
	sum := NewConvertableMoney(m[0]).Convert(currencyCode)
	for i := 1; i < len(m); i++ {
		convertedMoney := NewConvertableMoney(m[i]).Convert(currencyCode)
		sum, _ = sum.Add(convertedMoney)
	}
	return sum
}

func getExchangeRate(currencyCode, targetCurrencyCode string) decimal.Decimal {
	if currencyCode == "SGD" && targetCurrencyCode == "CNY" {
		return decimal.NewFromFloat(5.1)
	}

	if currencyCode == "CNY" && targetCurrencyCode == "CNY" {
		return decimal.NewFromFloat(1)
	}

	if currencyCode == "USD" && targetCurrencyCode == "CNY" {
		return decimal.NewFromFloat(7.1)
	}

	return decimal.NewFromFloat(1)
}
