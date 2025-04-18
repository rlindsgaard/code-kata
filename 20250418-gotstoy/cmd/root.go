/*
Copyright © 2025 Ronni Elken Lindsgaard <ronni dot lindsgaard at gmail dot com>
*/
package cmd

import (
	"os"

	"github.com/rlindsgaard/code-kata/20250418-gotstoy/config"
	"github.com/rlindsgaard/code-kata/20250418-gotstoy/input"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

var targetScope config.Scope
var targetEnsure config.Ensure
var updateAutomatically bool
var updateFrequency config.Frequency
var inputJSON *config.Settings

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotstoy",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(args []string) {
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Var(
		enumflag.New(&targetScope, "scope", config.ScopeMap, enumflag.EnumCaseInsensitive),
		"scope",
		"The target scope of the configuration",
	)
	rootCmd.RegisterFlagCompletionFunc("scope", config.ScopeFlagCompletion)

	rootCmd.PersistentFlags().Var(
		enumflag.New(&targetEnsure, "ensure", config.EnsureMap, enumflag.EnumCaseInsensitive),
		"ensure",
		"Whether the configuration should exist.",
	)
	rootCmd.RegisterFlagCompletionFunc("ensure", config.EnsureFlagCompletion)

	rootCmd.PersistentFlags().BoolVar(
		&updateAutomatically,
		"updateAutomatically",
		false,
		"Whether the configuration should set the app to update automatically.",
	)

	rootCmd.PersistentFlags().Var(
		&updateFrequency,
		"updateFrequency",
		"How often the app should check for updates between 1 and 90 days inclusive.",
	)

	rootCmd.PersistentFlags().Var(
		&input.JSONFlag{Target: &inputJSON},
		"inputJSON",
		"JSON string to be parsed",
	)
}
