package solvers

import (
	"fmt"
	"strings"
)

type Day7Solver struct{}

func init() {
	RegisterSolver(7, &Day7Solver{})
}

func (d *Day7Solver) Part1(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	positions := make(map[int]any)
	positions[strings.Index(lines[0], "S")] = nil
	splitCount := 0

	for _, line := range lines {
		newPositions := make(map[int]any)
		for position := range positions {
			if line[position] == '^' {
				splitCount++
				newPositions[position-1] = nil
				newPositions[position+1] = nil
			} else {
				newPositions[position] = nil
			}
		}
		positions = newPositions
	}

	return fmt.Sprintf("%d", splitCount), nil
}

func (d *Day7Solver) Part2(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	// Keep track of the number of times we can land at a position
	positions := make(map[int]int)
	positions[strings.Index(lines[0], "S")] = 1

	for _, line := range lines {
		newPositions := make(map[int]int)
		for position, val := range positions {
			if line[position] == '^' {
				newPositions[position-1] += val
				newPositions[position+1] += val
			} else {
				newPositions[position] += val
			}
		}
		positions = newPositions
	}

	sum := 0
	for _, val := range positions {
		sum += val
	}

	return fmt.Sprintf("%d", sum), nil
}
