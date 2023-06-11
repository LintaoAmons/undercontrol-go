/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/undercontrol/src/facade/cli"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new account",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Add()
	},
}

func init() {
	accountCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//  addCmd.Flags().StringP("name", "n", "Unnamed", "The name of this account")
	// addCmd.Flags().IntP("amount", "a", 0, "The amount of your account")
	// addCmd.Flags().StringP("Currency", "c", "CNY", "The currenty of this account")
}
