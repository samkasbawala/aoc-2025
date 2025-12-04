package solvers

import "testing"

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestDay1Part1(t *testing.T) {
	solver := Day1Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    testInput,
			expected: "3",
		},
		{
			name:     "land on 0",
			input:    "L50\n",
			expected: "1",
		},
		{
			name:     "land on 0 with no op",
			input:    "L50\nL0",
			expected: "1",
		},
		{
			name:     "land on 0, pass 0 > 1 time",
			input:    "L50\nR300",
			expected: "2",
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

func TestDay1Part2(t *testing.T) {
	solver := Day1Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    testInput,
			expected: "6",
		},
		{
			name:     "land on 0",
			input:    "L50\n",
			expected: "1",
		},
		{
			name:     "land on 0 with no op",
			input:    "L50\nL0",
			expected: "1",
		},
		{
			name:     "land on 0, pass 0 > 1 time",
			input:    "L50\nR300",
			expected: "4",
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
