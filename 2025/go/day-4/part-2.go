package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_PATH_P2 = "../inputs/day-4/part-1/input.txt"
)

func checkNeighborsP2(lines []string, i, j int) (neighborCount int) {
	min_i := 0
	max_i := len(lines) - 1
	min_j := 0
	neighborCount = 0

	for y := i - 1; y <= i+1; y++ {
		if y < min_i || y > max_i {
			continue
		}
		max_j_y := len(lines[y]) - 1
		for x := j - 1; x <= j+1; x++ {
			if x < min_j || x > max_j_y {
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

func checkLineP2(lines []string) (int, []string) {
	canAccessCount := 0
	copy := make([]string, len(lines))
	for i := range lines {
		copy[i] = lines[i]
	}

	toRemove := make(map[[2]int]bool)
	for i := range lines {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '@' {
				continue
			}
			neighborsCount := checkNeighborsP2(lines, i, j)
			if neighborsCount < 4 {
				toRemove[[2]int{i, j}] = true
				canAccessCount++
			}
		}
	}

	for pos := range toRemove {
		i, j := pos[0], pos[1]
		lineBytes := []byte(copy[i])
		lineBytes[j] = 'x'
		copy[i] = string(lineBytes)
	}

	return canAccessCount, copy
}

func Day4Part2() {
	file, err := os.Open(FILE_PATH_P2)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	var filteredLines []string
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	lines = filteredLines

	totalRemoved := 0
	for {
		validCount, newLines := checkLineP2(lines)
		if validCount == 0 {
			break
		}
		totalRemoved += validCount
		lines = newLines
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", totalRemoved)
}
