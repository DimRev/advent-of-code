package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_PATH_P2 = "../inputs/day-7/part-1/input.txt"
)

// PathCounter stores the number of paths at each position
type PathCounter []int

func generatePathCounter(line string) PathCounter {
	result := make([]int, len(line))
	for i, r := range line {
		if r == 'S' {
			result[i] = 1
		}
	}
	return result
}

func progressSplittersDP(line string, pathCounter PathCounter) PathCounter {
	newPathCounter := make([]int, len(line))

	for i, r := range line {
		if pathCounter[i] == 0 {
			continue
		}

		if r == '^' {
			if i-1 >= 0 {
				newPathCounter[i-1] += pathCounter[i]
			}
			if i+1 < len(line) {
				newPathCounter[i+1] += pathCounter[i]
			}
		} else {
			newPathCounter[i] += pathCounter[i]
		}
	}

	return newPathCounter
}

func Day7Part2() {
	file, err := os.Open(FILE_PATH_P2)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pathCounter := PathCounter{}
	idx := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if idx == 0 {
			pathCounter = generatePathCounter(line)
		} else if idx%2 == 0 {
			pathCounter = progressSplittersDP(line, pathCounter)
		}
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := 0
	for _, count := range pathCounter {
		solution += count
	}

	fmt.Printf("Solution: %d\n", solution)
}
