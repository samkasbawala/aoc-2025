package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/samkasbawala/aoc-2025/solvers"
	"github.com/spf13/cobra"
)

var part int
var inputFile string

func newDayCmd(dayNum int) *cobra.Command {
	dayCmd := &cobra.Command{
		Use:   fmt.Sprintf("day%0*d", 2, dayNum),
		Short: fmt.Sprintf("Solve day %0*d puzzle", 2, dayNum),
		Long:  fmt.Sprintf("Solve Advent of Code 2025 day %d puzzle", dayNum),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDay(dayNum)
		},
	}

	dayCmd.Flags().IntVarP(&part, "part", "p", 1, "Part of the puzzle to solve (1 or 2)")
	dayCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to input file (required)")
	dayCmd.MarkFlagRequired("input")

	return dayCmd
}

func runDay(dayNum int) error {
	if part != 1 && part != 2 {
		return fmt.Errorf("part must be 1 or 2, got %d", part)
	}

	// Read input file
	data, err := readInputFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Get the solver for this day
	solver := solvers.GetSolver(dayNum)
	if solver == nil {
		return fmt.Errorf("solver for day %d not yet implemented", dayNum)
	}

	// Call the appropriate solver function
	var result string
	if part == 1 {
		result, err = solver.Part1(data)
	} else {
		result, err = solver.Part2(data)
	}

	if err != nil {
		return fmt.Errorf("failed to solve day %d part %d: %w", dayNum, part, err)
	}

	fmt.Printf("Day %d, Part %d result: %s\n", dayNum, part, result)
	return nil
}

func readInputFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func registerDayCommands() {
	for i := 1; i <= 12; i++ {
		rootCmd.AddCommand(newDayCmd(i))
	}
}
