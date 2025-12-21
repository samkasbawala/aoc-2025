package solvers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Day3Solver struct{}

func init() {
	RegisterSolver(3, &Day3Solver{})
}

func (d *Day3Solver) Part1(input []byte) (string, error) {
	return solveWithDigits(input, 2)
}

func (d *Day3Solver) Part2(input []byte) (string, error) {
	return solveWithDigits(input, 12)
}

// findLargestJoltage finds the largest joltage in the bank recursively
func findLargestJoltage(bank string, digits int) (int64, error) {
	if digits <= 0 {
		return 0, nil
	}

	validRange := bank[0 : len(bank)-(digits-1)]
	maxSeen, maxIndex, err := findLargesNumber(validRange)
	if err != nil {
		return 0, fmt.Errorf("failed to find largest number: %v", err)
	}

	nextJoltage, err := findLargestJoltage(bank[maxIndex+1:], digits-1)
	if err != nil {
		return 0, fmt.Errorf("failed to find next joltage: %v", err)
	}

	return int64(maxSeen)*int64(math.Pow10(digits-1)) + nextJoltage, nil
}

// findLargestNumber simply finds the first largest number in the bank
func findLargesNumber(bank string) (int, int, error) {

	maxSeen := 0
	maxIndex := 0

	for index, c := range bank {
		numString := string(c)
		num, err := strconv.Atoi(numString)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to convert number to int: %v", err)
		}
		if num > maxSeen {
			maxSeen = num
			maxIndex = index
		}
	}

	return maxSeen, maxIndex, nil
}

// solveWithDigits processes all lines and finds the largest joltage with the given number of digits
func solveWithDigits(input []byte, digits int) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	wg := sync.WaitGroup{}
	results := make(chan int64, len(lines))

	for _, line := range lines {
		wg.Add(1)
		go func(bank string) {
			defer wg.Done()
			joltage, err := findLargestJoltage(bank, digits)
			if err != nil {
				results <- 0
				return
			}
			results <- joltage
		}(line)
	}
	wg.Wait()
	close(results)

	sum := int64(0)
	for result := range results {
		sum += result
	}

	return fmt.Sprintf("%d", sum), nil
}
