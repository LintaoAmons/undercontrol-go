package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LintaoAmons/undercontrol/src/usecase"
	"github.com/pterm/pterm"
)

var u = usecase.NewAccountUsecase()

func List() {
	accounts := u.FindAll()
	rows := [][]string{}
	rows = append(rows, []string{"Name", "Amount", "Currency"})
	for _, v := range accounts {
		rows = append(rows, []string{v.Name, fmt.Sprint(v.Amount.Display()), v.Amount.Currency().Code})
	}
	pterm.DefaultTable.WithHasHeader().WithData(rows).Render()
  // TODO: display total; display percentage
}

func Add() {
	name, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Your new account name")
	amountStr, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Amount of this account(0)")
	if amountStr == "" {
		amountStr = "0"
	}
	amount, err := strconv.ParseFloat(amountStr, 64)

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

func Get(name *string) {
	var nameInput string
	if name == nil {
		nameInput, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Name")
	} else {
		nameInput = *name
	}
	a, _ := u.Get(nameInput)
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)
	accountInfo := map[string]any{
		"Name":   a.Name,
		"Amount": a.Amount.Absolute().Display(),
	}
	logger.Info("========== Account Info ===========",
		logger.ArgsFromMap(accountInfo))
}

func Deposit() {
	accounts := u.FindAll()
	var options []string
	for _, v := range accounts {
		options = append(options, fmt.Sprint(v.Name+": "+v.Amount.Display()))
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select the account you want put money in")
	selectedName := strings.Split(selectedOption, ":")[0]

	amountStr, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Amount of this account(0)")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		pterm.Info.Printfln("Invalid amount")
	}
	u.Deposit(&usecase.DepositCommand{
		Name:   selectedName,
		Amount: amount,
	})
	after, _ := u.Get(selectedName)
	pterm.Info.Println("Deposit successfully")
	pterm.Info.Printfln("Account [%s]: %s", after.Name, after.Amount.Display())
}

func Withdraw() {
	accounts := u.FindAll()
	var options []string
	for _, v := range accounts {
		options = append(options, fmt.Sprint(v.Name+": "+v.Amount.Display()))
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select the account you want withdraw money out")
	selectedName := strings.Split(selectedOption, ":")[0]

	amountStr, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		pterm.Info.Printfln("Invalid amount")
	}
	u.Withdarw(&usecase.WithdarwCommand{
		Name:   selectedName,
		Amount: amount,
	})
	after, _ := u.Get(selectedName)
	pterm.Info.Println("Withdarw successfully")
	pterm.Info.Printfln("Account [%s]: %s", after.Name, after.Amount.Display())
}
