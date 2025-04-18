/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rlindsgaard/code-kata/20250418-gotstoy/config"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a tstoy configuration file to the desired state.",
	Long: `The set command ensures that the tstoy configuration file for a
specific scope has the desired settings. It returns the updated settings state
as a JSON blob to stdout.`,
	RunE: setState,
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setState(cmd *cobra.Command, args []string) error {
	enforcing := config.Settings{}

	if inputJSON != nil {
		enforcing = *inputJSON
	}
	if targetScope != config.ScopeUndefined {
		enforcing.Scope = targetScope
	}
	if targetEnsure != config.EnsureUndefined {
		enforcing.Ensure = targetEnsure
	}
	if rootCmd.PersistentFlags().Lookup("updateAutomatically").Changed {
		enforcing.UpdateAutomatically = &updateAutomatically
	}
	if updateFrequency != 0 {
		enforcing.UpdateFrequency = updateFrequency
	}

	final, err := enforcing.Enforce()

	if err != nil {
		return fmt.Errorf("can't enforce settings: %s", err)
	}
	return final.Print()
}
