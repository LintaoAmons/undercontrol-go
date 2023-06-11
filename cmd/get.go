/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/undercontrol/src/domain/usecase"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an account by account name",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			name, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Name")
		}
		usecase := usecase.NewAccountUsecase()
		a, _ := usecase.Get(name)
		logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)
		accountInfo := map[string]any{
			"Name":   a.Name,
			"Amount": a.Amount.Absolute().Display(),
		}
		logger.Info("========== Account Info ===========",
			logger.ArgsFromMap(accountInfo))
	},
}

func init() {
	accountCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}