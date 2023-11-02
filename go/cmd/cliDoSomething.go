/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lucaskatayama/learn-hexagonal/go/ui/cli"

	"github.com/spf13/cobra"
)

// cliDoSomethingCmd represents the cliDoSomething command
var cliDoSomethingCmd = &cobra.Command{
	Use:   "do-something",
	Short: "Do something through CLI",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		age, _ := cmd.Flags().GetInt("age")
		cli.DoSomething(cmd.Context(), name, age)
	},
}

func init() {
	cliCmd.AddCommand(cliDoSomethingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliDoSomethingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	cliDoSomethingCmd.Flags().StringP("name", "n", "John Doe", "Name")
	cliDoSomethingCmd.Flags().IntP("age", "a", 42, "Age")
}
