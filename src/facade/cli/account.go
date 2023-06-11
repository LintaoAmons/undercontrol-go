package cli

import (
	"fmt"

	"github.com/LintaoAmons/undercontrol/src/persistence/account"
	"github.com/pterm/pterm"
)

func List() {
	repo := account.NewAccountRepository()
	accounts := repo.FindAll()
	rows := [][]string{}
	rows = append(rows, []string{"Name", "Amount", "Currency"})
	for _, v := range accounts {
		rows = append(rows, []string{v.Name, fmt.Sprint(v.Amount.Absolute().Amount()), v.Amount.Currency().Code})
	}
	pterm.DefaultTable.WithHasHeader().WithData(rows).Render()
}
