/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/undercontrol/src/facade/cli"
	"github.com/spf13/cobra"
)

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit more money into one account",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Deposit()
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
