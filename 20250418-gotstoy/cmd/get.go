/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rlindsgaard/code-kata/20250418-gotstoy/config"
	"github.com/spf13/cobra"
)

var all bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the current state of a tstoy configuration file.",
	Long: `The get command returns the current state of a tstoy configuration
file as a JSON blob to stdout.`,
	RunE: getState,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVar(
		&all,
		"all",
		false,
		"Get configurations for all scopes.",
	)
}

func getState(cmd *cobra.Command, args []string) error {
	// Only the scope is used when retrieving current state.
	list := []config.Settings{}
	if all {
		list = append(
			list,
			config.Settings{Scope: config.ScopeMachine},
			config.Settings{Scope: config.ScopeUser},
		)
	} else if targetScope != config.ScopeUndefined {
		list = append(list, config.Settings{Scope: targetScope})
	} else if inputJSON != nil {
		list = append(list, *inputJSON)
	} else {
		list = append(list, config.Settings{Scope: targetScope})
	}

	for _, s := range list {
		err := s.Validate()
		if err != nil {
			return fmt.Errorf("can't get settings; %s", err)
		}

		config, err := s.GetConfigSettings()
		if err != nil {
			return fmt.Errorf("failed to get settings; %s", err)
		}

		err = config.Print()
		if err != nil {
			return err
		}
	}
	return nil
}
