package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type Day6Solver struct{}

func init() {
	RegisterSolver(6, &Day6Solver{})
}

func (d *Day6Solver) Part1(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	numbers := lines[0 : len(lines)-1]
	operators := lines[len(lines)-1]

	// Each item in the numbers slice is a string, where each number is separated by variable whitespace
	// Parse the numbers into a 2D slice of integers
	numbersInt := make([][]int64, len(numbers))
	for i, line := range numbers {
		numbersInt[i] = make([]int64, 0)
		numStrings := strings.Fields(line)

		for _, numString := range numStrings {
			numInt, err := strconv.ParseInt(numString, 10, 64)
			if err != nil {
				return "", fmt.Errorf("failed to convert number to int: %v", err)
			}
			numbersInt[i] = append(numbersInt[i], numInt)
		}
	}

	// Parse the operators
	operatorsSlice := strings.Fields(operators)

	// Transpose to make the operations easier
	transposedNumbers := transpose(numbersInt)

	result := int64(0)
	for i, operator := range operatorsSlice {
		switch operator {
		case "+":
			sum := int64(0)
			for _, num := range transposedNumbers[i] {
				sum += num
			}
			result += sum
		case "*":
			product := int64(1)
			for _, num := range transposedNumbers[i] {
				product *= num
			}
			result += product
		}
	}

	return fmt.Sprintf("%d", result), nil
}

func transpose[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 {
		return [][]T{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]T, cols)
	for i := range transposed {
		transposed[i] = make([]T, rows)
	}

	for i := range rows {
		for j := range cols {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func (d *Day6Solver) Part2(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	numbers := lines[0 : len(lines)-1]
	operators := lines[len(lines)-1]
	operatorsSlice := strings.Fields(operators)

	result := int64(0)
	var current int64
	currentOperatorIndex := 0

	switch operatorsSlice[currentOperatorIndex] {
	case "+":
		current = 0
	case "*":
		current = 1
	}

	for col := range len(numbers[0]) {
		num := ""
		for row := range len(numbers) {
			if string(numbers[row][col]) != " " {
				num += string(numbers[row][col])
			}
		}

		// Moving on to next operator
		if num == "" {
			result += current

			currentOperatorIndex++
			switch operatorsSlice[currentOperatorIndex] {
			case "+":
				current = 0
			case "*":
				current = 1
			}
			continue
		}

		numInt, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			return "", fmt.Errorf("failed to convert number to int: %v", err)
		}

		switch operatorsSlice[currentOperatorIndex] {
		case "+":
			current += numInt
		case "*":
			current *= numInt
		}
	}

	// Need to add the last current value
	result += current

	return fmt.Sprintf("%d", result), nil
}
