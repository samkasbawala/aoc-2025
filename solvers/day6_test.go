package solvers

import (
	"testing"
)

const day6TestInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestDay6Part1(t *testing.T) {
	solver := Day6Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day6TestInput,
			expected: "4277556",
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

func TestDay6Part2(t *testing.T) {
	solver := Day6Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day6TestInput,
			expected: "3263827",
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
