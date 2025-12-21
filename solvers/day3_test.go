package solvers

import (
	"testing"
)

const day3TestInput = `987654321111111
811111111111119
234234234234278
818181911112111
`

func TestDay3FindLargestNumber(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValue int
		expectedIndex int
		expectedError bool
	}{
		{
			name:          "largest number at the beginning",
			input:         "987654321111111",
			expectedValue: 9,
			expectedIndex: 0,
			expectedError: false,
		},
		{
			name:          "largest number at the end",
			input:         "811111111111119",
			expectedValue: 9,
			expectedIndex: 14,
			expectedError: false,
		},
		{
			name:          "largest number in the middle",
			input:         "12344321",
			expectedValue: 4,
			expectedIndex: 3,
			expectedError: false,
		},
		{
			name:          "non integer input",
			input:         "1234567890a",
			expectedValue: 0,
			expectedIndex: 0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotIndex, gotError := findLargesNumber(tt.input)
			if (gotError != nil) != tt.expectedError {
				t.Errorf("got error %v, want %v", gotError, tt.expectedError)
			}

			if gotValue != tt.expectedValue {
				t.Errorf("got value %d, want %d", gotValue, tt.expectedValue)
			}

			if gotIndex != tt.expectedIndex {
				t.Errorf("got index %d, want %d", gotIndex, tt.expectedIndex)
			}
		})
	}
}

func TestDay3FindLargestJoltage(t *testing.T) {
	tests := []struct {
		name          string
		inputBank     string
		inputDigits   int
		expectedValue int64
		expectedError bool
	}{
		{
			name:          "sample 1",
			inputBank:     "987654321111111",
			inputDigits:   2,
			expectedValue: 98,
			expectedError: false,
		},
		{
			name:          "sample 2",
			inputBank:     "811111111111119",
			inputDigits:   2,
			expectedValue: 89,
			expectedError: false,
		},
		{
			name:          "sample 3",
			inputBank:     "234234234234278",
			inputDigits:   2,
			expectedValue: 78,
			expectedError: false,
		},
		{
			name:          "sample 4",
			inputBank:     "818181911112111",
			inputDigits:   2,
			expectedValue: 92,
			expectedError: false,
		},
		{
			name:          "sample 5",
			inputBank:     "987654321111111",
			inputDigits:   12,
			expectedValue: 987654321111,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotError := findLargestJoltage(tt.inputBank, tt.inputDigits)

			if (gotError != nil) != tt.expectedError {
				t.Errorf("got error %v, want %v", gotError, tt.expectedError)
			}

			if gotValue != tt.expectedValue {
				t.Errorf("got value %d, want %d", gotValue, tt.expectedValue)
			}
		})
	}
}

func TestDay3Part1(t *testing.T) {
	solver := Day3Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day3TestInput,
			expected: "357",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solver.Part1([]byte(tt.input))

			if err != nil {
				t.Fatalf("error: err %s", err)
			}

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}

func TestDay3Part2(t *testing.T) {
	solver := Day3Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day3TestInput,
			expected: "3121910778619",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solver.Part2([]byte(tt.input))

			if err != nil {
				t.Fatalf("error: err %s", err)
			}

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}
