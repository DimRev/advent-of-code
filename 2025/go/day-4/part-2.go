package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_PATH_P2 = "../inputs/day-4/part-1/temp.txt"
)

func checkNeighborsP2(lines []string, i, j int) (neighborCount int) {
	min_i := 0
	max_i := len(lines) - 1
	min_j := 0
	max_j := len(lines[0]) - 1
	neighborCount = 0

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

			if lines[y][x] == '@' || lines[y][x] == 'x' {
				neighborCount++
			}

		}
	}

	return neighborCount
}

func checkLineP2(lines []string, checkIdx int) (int, string) {
	canAccessCount := 0
	for i := 0; i < len(lines); i++ {
		if i != checkIdx {
			continue
		}
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '@' {
				continue
			}
			neighborsCount := checkNeighborsP2(lines, i, j)

			if neighborsCount < 4 {
				lineBytes := []byte(lines[i])
				lineBytes[j] = 'x'
				lines[i] = string(lineBytes)
				canAccessCount++
			}
		}
	}

	fmt.Printf("%v\n", lines[checkIdx])

	return canAccessCount, lines[checkIdx]
}

func Day4Part2() {
	file, err := os.Open(FILE_PATH_P2)
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
			c, _ := checkLineP2(lines, 0)
			validCount += c
		}

		if len(lines) == 3 {
			c, _ := checkLineP2(lines, 1)
			validCount += c
			lines = lines[1:]
		}
	}

	c, _ := checkLineP2(lines, 1)
	validCount += c

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", validCount)
}
