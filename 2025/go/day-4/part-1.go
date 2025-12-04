package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-4/part-1/temp.txt"
)

func checkNeighbors(lines []string, i, j int) int {
	min_i := 0
	max_i := len(lines) - 1
	min_j := 0
	max_j := len(lines[0]) - 1
	neighborCount := 0

	for y := i - 1; y <= i+1; y++ {
		if y < min_i || y > max_i {
			continue
		}
		for x := j - 1; x <= j+1; x++ {
			if x < min_j || x > max_j {
				continue
			}
			if y == i && x == j {
				continue
			}

			if lines[y][x] == '@' {
				neighborCount++
			}

		}
	}

	return neighborCount
}

func checkLine(lines []string, checkIdx int) int {
	canAccessCount := 0
	for i := 0; i < len(lines); i++ {
		if i != checkIdx {
			continue
		}
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '@' {
				continue
			}
			neighborsCount := checkNeighbors(lines, i, j)
			if neighborsCount < 4 {
				canAccessCount++
			}
		}
	}
	return canAccessCount
}

func Day4Part1() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validCount := 0
	lines := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
		if len(lines) == 2 {
			validCount += checkLine(lines, 0)
		}

		if len(lines) == 3 {
			validCount += checkLine(lines, 1)
			lines = lines[1:]
		}
	}

	validCount += checkLine(lines, 1)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", validCount)
}
