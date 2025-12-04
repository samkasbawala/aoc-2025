package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-2025",
	Short: "Advent of Code 2025 solutions",
	Long:  "A CLI application for solving Advent of Code 2025 puzzles in Go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Register day commands
	registerDayCommands()
}

