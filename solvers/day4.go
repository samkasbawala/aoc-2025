package solvers

import (
	"fmt"
	"strings"
)

type Day4Solver struct{}

func init() {
	RegisterSolver(4, &Day4Solver{})
}

func (d *Day4Solver) Part1(input []byte) (string, error) {
	grid, dy, dx := createGrid(input)

	count := countAccessible(grid, dy, dx, false)

	return fmt.Sprintf("%d", count), nil
}

func (d *Day4Solver) Part2(input []byte) (string, error) {
	grid, dy, dx := createGrid(input)

	count := 0
	for {
		removed := countAccessible(grid, dy, dx, true)
		if removed == 0 {
			break
		}
		count += removed
	}

	return fmt.Sprintf("%d", count), nil
}

// createGrid parses the input and creates a 2D grid along with its dimensions
func createGrid(input []byte) ([][]string, int, int) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	dy := len(lines)
	dx := len(lines[0])

	grid := make([][]string, dy)
	for i := range grid {
		grid[i] = make([]string, dx)
	}

	for y, line := range lines {
		for x, char := range line {
			grid[y][x] = string(char)
		}
	}

	return grid, dy, dx
}

// countAccessible counts the number of accessible positions in the grid
func countAccessible(grid [][]string, dy int, dx int, remove bool) int {

	count := 0

	for y := range dy {
		for x := range dx {
			if grid[y][x] == "@" && isAccessible(grid, dy, dx, y, x, remove) {
				count++
			}
		}
	}

	return count
}

// isAccessible checks if the current position is accessible.
// The current position is accessible if there are fewer than 4 "@" in the eight adjacent positions.
func isAccessible(grid [][]string, dy, dx, y, x int, remove bool) bool {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			// Check if the position is within the grid
			if y+i < 0 || y+i >= dy || x+j < 0 || x+j >= dx {
				continue
			}

			if grid[y+i][x+j] == "@" {
				count++
			}
		}
	}

	if count >= 4 {
		return false
	} else {
		if remove {
			grid[y][x] = "x"
		}
		return true
	}
}
