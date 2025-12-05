package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2Solver struct{}

func init() {
	RegisterSolver(2, &Day2Solver{})
}

func (d Day2Solver) Part1(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	str := lines[0]
	ranges := strings.Split(strings.TrimSpace(str), ",")

	sumInvalid := 0

	for _, numRange := range ranges {
		nums := strings.Split(numRange, "-")
		lower := nums[0]
		upper := nums[1]

		lowerInt, err := strconv.Atoi(lower)
		if err != nil {
			return "", fmt.Errorf("failed to convert lower to int: %v", err)
		}
		upperInt, err := strconv.Atoi(upper)
		if err != nil {
			return "", fmt.Errorf("failed to convert upper to int: %v", err)
		}

		for i := lowerInt; i <= upperInt; i++ {
			numStr := strconv.Itoa(i)

			// Only need to check if the number of digits is even
			if len(numStr)%2 == 0 {
				halfLen := len(numStr) / 2
				if numStr[:halfLen] == numStr[halfLen:] {
					sumInvalid += i
				}
			}
		}

	}

	return fmt.Sprintf("%d", sumInvalid), nil
}

func (d Day2Solver) Part2(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	str := lines[0]
	ranges := strings.Split(strings.TrimSpace(str), ",")

	sumInvalid := 0

	for _, numRange := range ranges {
		nums := strings.Split(numRange, "-")
		lower := nums[0]
		upper := nums[1]

		lowerInt, err := strconv.Atoi(lower)
		if err != nil {
			return "", fmt.Errorf("failed to convert lower to int: %v", err)
		}
		upperInt, err := strconv.Atoi(upper)
		if err != nil {
			return "", fmt.Errorf("failed to convert upper to int: %v", err)
		}

		for i := lowerInt; i <= upperInt; i++ {
			numStr := strconv.Itoa(i)

			// Check each substring that is no longer than half the length of the number string
			for j := 1; j <= len(numStr)/2; j++ {

				// j must be a multiple of len(numStr)
				if len(numStr)%j != 0 {
					continue
				}

				// Must occur at least twice
				multiple := len(numStr) / j
				if multiple < 2 {
					continue
				}
				repeated := strings.Repeat(numStr[:j], multiple)

				// Break since we've found an invalid number and no need to check further
				if repeated == numStr {
					sumInvalid += i
					break
				}
			}
		}

	}

	return fmt.Sprintf("%d", sumInvalid), nil
}
