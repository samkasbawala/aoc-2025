package solvers

import (
	"testing"
)

const day5TestInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestDay5Part1(t *testing.T) {
	solver := Day5Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day5TestInput,
			expected: "3",
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

func TestDay5Part2(t *testing.T) {
	solver := Day5Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day5TestInput,
			expected: "14",
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
