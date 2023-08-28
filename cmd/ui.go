/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LintaoAmons/undercontrol/src/facade/ui"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Run as TUI",
	Run: func(cmd *cobra.Command, args []string) {
		ui.InitTui()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
