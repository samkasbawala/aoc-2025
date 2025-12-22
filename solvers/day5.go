package solvers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day5Solver struct{}

func init() {
	RegisterSolver(5, &Day5Solver{})
}

type freshRange struct {
	lower int64
	upper int64
}

func (d *Day5Solver) Part1(input []byte) (string, error) {
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rangesRaw := strings.Split(parts[0], "\n")
	ingredients := strings.Split(parts[1], "\n")

	// Parse the ranges
	freshRanges, err := constructFreshRanges(rangesRaw)
	if err != nil {
		return "", fmt.Errorf("failed to construct fresh ranges: %v", err)
	}

	count := 0

	for _, ingredient := range ingredients {
		ingredientInt, err := strconv.ParseInt(ingredient, 10, 64)
		if err != nil {
			return "", fmt.Errorf("failed to convert ingredient to int: %v", err)
		}

		// Find the first range where lower > ingredientInt (binary search)
		idx := sort.Search(len(freshRanges), func(i int) bool {
			return freshRanges[i].lower > ingredientInt
		})

		// Check the previous range (if any) to see if it contains the ingredient
		if idx > 0 {
			prevIdx := idx - 1
			if freshRanges[prevIdx].upper >= ingredientInt {
				count++
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

// constructFreshRanges constructs the fresh ranges from the raw ranges text input.
// The ranges will be sorted and then merged to create the final fresh ranges.
func constructFreshRanges(rangesRaw []string) ([]freshRange, error) {

	freshRanges := make([]freshRange, 0)

	// Parse the ranges
	for _, r := range rangesRaw {
		bounds := strings.Split(r, "-")
		lower, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert lower to int: %v", err)
		}

		upper, err := strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert upper to int: %v", err)
		}

		freshRanges = append(freshRanges, freshRange{lower, upper})
	}

	// Sort the ranges by the lower bound, then the upper bound
	sort.SliceStable(freshRanges, func(i, j int) bool {
		return freshRanges[i].lower < freshRanges[j].lower || (freshRanges[i].lower == freshRanges[j].lower && freshRanges[i].upper < freshRanges[j].upper)
	})

	// Merge the ranges
	i := 0
	j := 0
	for j < len(freshRanges) {
		if freshRanges[i].upper >= freshRanges[j].lower {
			freshRanges[i].upper = max(freshRanges[i].upper, freshRanges[j].upper)
			j++
		} else {
			freshRanges[i+1] = freshRanges[j]
			i++
			j++
		}
	}
	freshRanges = freshRanges[:i+1]

	return freshRanges, nil
}

func (d *Day5Solver) Part2(input []byte) (string, error) {
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rangesRaw := strings.Split(parts[0], "\n")

	freshRanges, err := constructFreshRanges(rangesRaw)
	if err != nil {
		return "", fmt.Errorf("failed to construct fresh ranges: %v", err)
	}

	count := int64(0)

	for _, freshRange := range freshRanges {
		count += freshRange.upper - freshRange.lower + 1
	}

	return fmt.Sprintf("%d", count), nil
}
