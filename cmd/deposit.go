/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LintaoAmons/undercontrol/src/usecase"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit more money into one account",
	Run: func(cmd *cobra.Command, args []string) {
		u := usecase.NewAccountUsecase()
		accounts := u.FindAll()
		var options []string
		for _, v := range accounts {
			options = append(options, fmt.Sprint(v.Name+": "+v.Amount.Display()))
		}
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select the account you want put money in")
		selectedName := strings.Split(selectedOption, ":")[0]

		amountStr, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Amount of this account(0)")
		pterm.Info.Printfln("selectedOption: %s", pterm.Green(selectedOption))
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			pterm.Info.Printfln("Invalid amount")
		}
		u.Deposit(&usecase.DepositCommand{
			Name:   selectedName,
			Amount: amount,
		})
	},
}

func init() {
	accountCmd.AddCommand(depositCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depositCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// depositCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
