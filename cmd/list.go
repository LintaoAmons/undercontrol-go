/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	account "github.com/LintaoAmons/undercontrol/src/persistence/account"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		repo := account.NewAccountRepository()
		accounts := repo.FindAll()
		rows := [][]string{}
    rows = append(rows, []string{"Name", "Amount", "Currency"})
		for _, v := range accounts {
      rows = append(rows, []string{v.Name, fmt.Sprint(v.Amount.Absolute().Amount()), v.Amount.Currency().Code})
		}
		pterm.DefaultTable.WithHasHeader().WithData(rows).Render()
	},
}

func init() {
	accountCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
