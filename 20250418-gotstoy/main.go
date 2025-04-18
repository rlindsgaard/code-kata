/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/rlindsgaard/code-kata/20250418-gotstoy/cmd"
	"github.com/rlindsgaard/code-kata/20250418-gotstoy/input"
)

func main() {
	args := []string{}
	for index, arg := range os.Args {
		if index > 0 {
			args = append(args, arg)
		}
	}
	args = input.HandleStdIn(args)

	cmd.Execute(args)
}
