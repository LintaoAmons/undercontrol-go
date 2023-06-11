package cli

import (
	"fmt"
	"strconv"

	"github.com/LintaoAmons/undercontrol/src/usecase"
	"github.com/pterm/pterm"
)

var u = usecase.NewAccountUsecase()

func List() {
	accounts := u.FindAll()
	rows := [][]string{}
	rows = append(rows, []string{"Name", "Amount", "Currency"})
	for _, v := range accounts {
		rows = append(rows, []string{v.Name, fmt.Sprint(v.Amount.Absolute().Amount()), v.Amount.Currency().Code})
	}
	pterm.DefaultTable.WithHasHeader().WithData(rows).Render()
}

func Add() {
	name, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Your new account name")
	amountStr, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Amount of this account(0)")
	if amountStr == "" {
		amountStr = "0"
	}
	amount, err := strconv.Atoi(amountStr)
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo).WithMaxWidth(200)
	if err != nil {
		logger.Error("Invalid amount", []pterm.LoggerArgument{
			{
				Key:   "Your input",
				Value: "[" + amountStr + "]",
			},
		})
		panic("Invalid Amount") // TODO: is there another way to exit the app other than panic?
	}
	currency, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Currency of this account(CNY)")
	if currency == "" {
		currency = "CNY"
	}

	u.Add(&usecase.CreateAccountCommand{
		Name:     name,
		Amount:   amount,
		Currency: currency,
	})
	logger.Info("Created successfully")
}
