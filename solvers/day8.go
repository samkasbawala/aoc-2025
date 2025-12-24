package solvers

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day8Solver struct{}

func init() {
	RegisterSolver(8, &Day8Solver{})
}

type junctionBox struct {
	x         float64
	y         float64
	z         float64
	neighbors map[*junctionBox]any
}

type junctionBoxDistance struct {
	junctionBox1 *junctionBox
	junctionBox2 *junctionBox
	distance     float64
}

func (j *junctionBox) straightLineDistance(other *junctionBox) float64 {
	distance := math.Pow(j.x-other.x, 2) + math.Pow(j.y-other.y, 2) + math.Pow(j.z-other.z, 2)
	return math.Sqrt(distance)
}

func (d *Day8Solver) Part1(input []byte) (string, error) {

	junctionBoxes, err := createJunctionBoxes(input)
	if err != nil {
		return "", fmt.Errorf("failed to create junction boxes: %v", err)
	}

	// Create a list of distances between all junction boxes
	junctionBoxesDistances := []*junctionBoxDistance{}
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			junctionBoxesDistances = append(junctionBoxesDistances, &junctionBoxDistance{junctionBoxes[i], junctionBoxes[j], junctionBoxes[i].straightLineDistance(junctionBoxes[j])})
		}
	}

	// Sort the distances in ascending order
	sort.SliceStable(junctionBoxesDistances, func(i, j int) bool {
		return junctionBoxesDistances[i].distance < junctionBoxesDistances[j].distance
	})

	// Construct the graph by connecting the closest 1000 junction boxes
	makeCircuits(junctionBoxesDistances, 1000)

	connectedComponents := []int{}
	visited := make(map[*junctionBox]any)
	for _, box := range junctionBoxes {
		if _, ok := visited[box]; !ok {
			connectedComponents = append(connectedComponents, dfs(box, visited))
		}
	}

	// Sort the connected components in descending order
	sort.SliceStable(connectedComponents, func(i, j int) bool {
		return connectedComponents[i] > connectedComponents[j]
	})

	// The answer is the sum of the 1000 largest connected components
	product := 1
	for i := range 3 {
		product *= connectedComponents[i]
	}

	return fmt.Sprintf("%d", product), nil
}

func createJunctionBoxes(input []byte) ([]*junctionBox, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	junctionBoxes := make([]*junctionBox, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert x to float64: %v", err)
		}
		y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert y to float64: %v", err)
		}
		z, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert z to float64: %v", err)
		}
		junctionBoxes = append(junctionBoxes, &junctionBox{x, y, z, map[*junctionBox]any{}})
	}

	return junctionBoxes, nil
}

func makeCircuits(junctionBoxesDistances []*junctionBoxDistance, connectionLimit int) {
	for i := range connectionLimit {
		box1 := junctionBoxesDistances[i].junctionBox1
		box2 := junctionBoxesDistances[i].junctionBox2

		box1.neighbors[box2] = nil
		box2.neighbors[box1] = nil
	}
}

func dfs(start *junctionBox, visited map[*junctionBox]any) int {
	visited[start] = nil
	count := 1
	for neighbor := range start.neighbors {
		if _, ok := visited[neighbor]; !ok {
			count += dfs(neighbor, visited)
		}
	}
	return count
}

func (d *Day8Solver) Part2(input []byte) (string, error) {
	junctionBoxes, err := createJunctionBoxes(input)
	if err != nil {
		return "", fmt.Errorf("failed to create junction boxes: %v", err)
	}

	// Create a list of distances between all junction boxes
	junctionBoxesDistances := []*junctionBoxDistance{}
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			junctionBoxesDistances = append(junctionBoxesDistances, &junctionBoxDistance{junctionBoxes[i], junctionBoxes[j], junctionBoxes[i].straightLineDistance(junctionBoxes[j])})
		}
	}

	// Sort the distances in ascending order
	sort.SliceStable(junctionBoxesDistances, func(i, j int) bool {
		return junctionBoxesDistances[i].distance < junctionBoxesDistances[j].distance
	})

	// "Add" the paths from shortest to longest, stopping when all junction boxes are connected
	seen := make(map[*junctionBox]any)
	var result int64
	for _, distance := range junctionBoxesDistances {

		if _, ok := seen[distance.junctionBox1]; !ok {
			seen[distance.junctionBox1] = nil
		}
		if _, ok := seen[distance.junctionBox2]; !ok {
			seen[distance.junctionBox2] = nil
		}

		if len(seen) == len(junctionBoxes) {
			result = int64(distance.junctionBox1.x * distance.junctionBox2.x)
			break
		}
	}

	return fmt.Sprintf("%d", result), nil
}
