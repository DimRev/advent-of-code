package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	circleSize    = 100
	startPosition = 50
)

func Day1Part1() {
	startTs := time.Now()
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

		switch direction {
		case "L":
			position -= distance
		case "R":
			position += distance
		default:
			fmt.Printf("Error parsing line: %v\n", err)
			os.Exit(1)
		}

		position = wrapPosition(position)

		if position == 0 {
			crossingsAtZero++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	endTs := time.Now()
	elapsed := endTs.Sub(startTs)
	fmt.Printf("Solution: %d [%d(μs)]\n", crossingsAtZero, elapsed.Microseconds())
}

func wrapPosition(position int) (wrappedPosition int) {
	return ((position % circleSize) + circleSize) % circleSize
}

func parseLine(line string) (direction string, distance int, err error) {
	if len(line) < 2 {
		return "", 0, fmt.Errorf("invalid line: %s", line)
	}

	distance, err = strconv.Atoi(line[1:])
	if err != nil {
		return "", 0, err
	}

	return line[:1], distance, nil
}
