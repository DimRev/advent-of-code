package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	n := 50
	z := 0
	for scanner.Scan() {
		dir, dist, err := parsePartUnbounded(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Printf("Error parsing part: %v\n", err)
			os.Exit(1)
		}

		fullRotationsFromDist := dist / 100
		z += fullRotationsFromDist
		dist = dist % 100

		passesOverZero := calculatePassesOverZero(dir, n, dist)
		z += passesOverZero

		if dir == "L" {
			n -= dist
		} else {
			n += dist
		}
		n = ((n % 100) + 100) % 100

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", z)
}

func calculatePassesOverZero(dir string, startPos int, distance int) int {
	if distance == 0 {
		return 0
	}

	if dir == "R" {
		distanceToZero := 100 - startPos
		if startPos == 0 {
			distanceToZero = 100
		}
		if distance >= distanceToZero {
			return 1 + (distance-distanceToZero)/100
		}
		return 0
	} else {
		distanceToZero := startPos
		if startPos == 0 {
			distanceToZero = 100
		}
		if distance >= distanceToZero {
			return 1 + (distance-distanceToZero)/100
		}
		return 0
	}
}

func parsePartUnbounded(part string) (dir string, dist int, err error) {
	if len(part) < 2 {
		return "", 0, fmt.Errorf("invalid part: %s", part)
	}

	dist, err = strconv.Atoi(part[1:])
	if err != nil {
		return "", 0, err
	}

	return part[:1], dist, nil
}
