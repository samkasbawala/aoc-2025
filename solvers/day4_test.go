package solvers

import (
	"testing"
)

const day4TestInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestDay4Part1(t *testing.T) {
	solver := Day4Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day4TestInput,
			expected: "13",
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

func TestDay4Part2(t *testing.T) {
	solver := Day4Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day4TestInput,
			expected: "43",
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
