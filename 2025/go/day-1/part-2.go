package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day1Part2() {
	file, err := os.Open("../inputs/day-1/part-1/input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := startPosition
	crossingsAtZero := 0

	for scanner.Scan() {
		direction, distance, err := parseLine(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Printf("Error parsing line: %v\n", err)
			os.Exit(1)
		}

		fullRotations := distance / circleSize
		crossingsAtZero += fullRotations

		remainingDistance := distance % circleSize
		additionalCrossings := calculatePassesOverZero(direction, position, remainingDistance)
		crossingsAtZero += additionalCrossings

		switch direction {
		case "L":
			position -= remainingDistance
		case "R":
			position += remainingDistance
		default:
			fmt.Printf("Error parsing line: %v\n", err)
		}

		position = wrapPosition(position)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", crossingsAtZero)
}

func calculatePassesOverZero(direction string, startPos int, distance int) (numberOfPassesOverZero int) {
	if distance == 0 {
		return 0
	}

	distanceToZero := 0
	if direction == "R" {
		if startPos == 0 {
			distanceToZero = circleSize
		} else {
			distanceToZero = circleSize - startPos
		}
	} else {
		if startPos == 0 {
			distanceToZero = circleSize
		} else {
			distanceToZero = startPos
		}
	}

	if distance >= distanceToZero {
		return 1 + (distance-distanceToZero)/circleSize
	}
	return 0
}
