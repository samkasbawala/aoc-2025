package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type Day1Solver struct{}

func init() {
	RegisterSolver(1, &Day1Solver{})
}

func (d *Day1Solver) Part1(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	position := 50
	counter := 0

	for _, line := range lines {
		direction := line[:1]
		distance := line[1:]

		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			return "", fmt.Errorf("failed to convert distance to int: %v", err)
		}

		// Do nothing if it's a no op
		if distanceInt <= 0 {
			continue
		}

		if strings.ToUpper(direction) == "L" {
			position = (position - distanceInt) % 100
		} else {
			position = (position + distanceInt) % 100
		}

		if position == 0 {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

func (d *Day1Solver) Part2(input []byte) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	position := 50
	counter := 0

	for _, line := range lines {
		direction := line[:1]
		distance := line[1:]

		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			return "", fmt.Errorf("failed to convert distance to int: %v", err)
		}

		counter += (distanceInt / 100)
		distanceInt = distanceInt % 100

		if distanceInt != 0 {
			prev := position
			if direction == "L" {
				position -= distanceInt
				if position <= 0 && 0 < prev {
					counter++
				}
			} else {
				position += distanceInt
				if position >= 100 {
					counter++
				}
			}

			if position < 0 {
				position += 100
			} else if position >= 100 {
				position -= 100
			}

		}
	}
	return fmt.Sprintf("%d", counter), nil
}
