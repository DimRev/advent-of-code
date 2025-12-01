package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1Part1() {
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
		dir, dist, err := parsePart(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Printf("Error parsing part: %v\n", err)
			os.Exit(1)
		}

		if dir == "L" {
			n -= dist
		} else {
			n += dist
		}

		n = ((n % 100) + 100) % 100
		if n == 0 {
			z++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", z)
}

func parsePart(part string) (dir string, dist int, err error) {
	if len(part) < 2 {
		return "", 0, fmt.Errorf("invalid part: %s", part)
	}

	dist, err = strconv.Atoi(part[1:])
	if err != nil {
		return "", 0, err
	}

	return part[:1], dist % 100, nil
}
