package solvers

import (
	"reflect"
	"sort"
	"testing"
)

const day8TestInput = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestCreateJunctionBoxes(t *testing.T) {
	junctionBoxes, err := createJunctionBoxes([]byte(day8TestInput))
	if err != nil {
		t.Fatalf("error: err %s", err)
	}

	if len(junctionBoxes) != 20 {
		t.Errorf("got %d, want %d", len(junctionBoxes), 10)
	}

	expected := []*junctionBox{
		{162, 817, 812, map[*junctionBox]any{}},
		{57, 618, 57, map[*junctionBox]any{}},
		{906, 360, 560, map[*junctionBox]any{}},
		{592, 479, 940, map[*junctionBox]any{}},
		{352, 342, 300, map[*junctionBox]any{}},
		{466, 668, 158, map[*junctionBox]any{}},
		{542, 29, 236, map[*junctionBox]any{}},
		{431, 825, 988, map[*junctionBox]any{}},
		{739, 650, 466, map[*junctionBox]any{}},
		{52, 470, 668, map[*junctionBox]any{}},
		{216, 146, 977, map[*junctionBox]any{}},
		{819, 987, 18, map[*junctionBox]any{}},
		{117, 168, 530, map[*junctionBox]any{}},
		{805, 96, 715, map[*junctionBox]any{}},
		{346, 949, 466, map[*junctionBox]any{}},
		{970, 615, 88, map[*junctionBox]any{}},
		{941, 993, 340, map[*junctionBox]any{}},
		{862, 61, 35, map[*junctionBox]any{}},
		{984, 92, 344, map[*junctionBox]any{}},
		{425, 690, 689, map[*junctionBox]any{}},
	}

	if !reflect.DeepEqual(junctionBoxes, expected) {
		t.Errorf("got %v, want %v", junctionBoxes, expected)
	}
}

func TestDay8Part1(t *testing.T) {
	junctionBoxes, err := createJunctionBoxes([]byte(day8TestInput))
	if err != nil {
		t.Fatalf("error: err %s", err)
	}

	junctionBoxesDistances := []*junctionBoxDistance{}
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			junctionBoxesDistances = append(junctionBoxesDistances, &junctionBoxDistance{junctionBoxes[i], junctionBoxes[j], junctionBoxes[i].straightLineDistance(junctionBoxes[j])})
		}
	}

	sort.SliceStable(junctionBoxesDistances, func(i, j int) bool {
		return junctionBoxesDistances[i].distance < junctionBoxesDistances[j].distance
	})

	makeCircuits(junctionBoxesDistances, 10)

	connectedComponents := []int{}
	visited := make(map[*junctionBox]any)
	for _, box := range junctionBoxes {
		if _, ok := visited[box]; !ok {
			connectedComponents = append(connectedComponents, dfs(box, visited))
		}
	}

	sort.SliceStable(connectedComponents, func(i, j int) bool {
		return connectedComponents[i] > connectedComponents[j]
	})

	product := 1
	for i := range 3 {
		product *= connectedComponents[i]
	}

	expected := 40

	if product != expected {
		t.Errorf("got %d, want %d", product, expected)
	}
}

func TestDay8Part2(t *testing.T) {
	solver := Day8Solver{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sample",
			input:    day8TestInput,
			expected: "25272",
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
