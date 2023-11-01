/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lucaskatayama/learn-hexagonal/go/ui/rest"

	"github.com/spf13/cobra"
)

// restRunCmd represents the restRun command
var restRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run REST server",
	Run: func(cmd *cobra.Command, args []string) {
		rest.Run()
	},
}

func init() {
	restCmd.AddCommand(restRunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
