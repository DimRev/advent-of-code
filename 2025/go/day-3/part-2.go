package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MAX_DIGITS = 12
)

func maxPairD2(line string) int {
	maxIdx := -1
	result := 0

	for digitPos := MAX_DIGITS - 1; digitPos >= 0; digitPos-- {
		maxDigit := -1

		for i := maxIdx + 1; i < len(line)-digitPos; i++ {
			num, err := parseDigit(line[i])
			if err != nil {
				fmt.Printf("Error parsing digit: %v\n", err)
				os.Exit(1)
			}
			if num > maxDigit {
				maxDigit = num
				maxIdx = i
			}
		}

		result = result*10 + maxDigit
	}

	return result
}

func Day3Part2() {
	file, err := os.Open("../inputs/day-3/part-1/input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pairSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if err != nil {
			fmt.Printf("Error parsing line: %v\n", err)
			os.Exit(1)
		}

		maxPair := maxPairD2(line)
		pairSum += maxPair
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", pairSum)
}
