package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-3/part-1/input.txt"
)

func parseDigit(b byte) (int, error) {
	return strconv.Atoi(string(b))
}

func maxPairD1(line string) int {
	firstMaxDigit := -1
	firstMaxDigitIndex := -1

	for i := 0; i < len(line)-1; i++ {
		num, err := parseDigit(line[i])
		if err != nil {
			fmt.Printf("Error parsing digit: %v\n", err)
			os.Exit(1)
		}
		if num > firstMaxDigit {
			firstMaxDigit = num
			firstMaxDigitIndex = i
		}
	}

	secondMaxDigit := -1

	for i := firstMaxDigitIndex + 1; i < len(line); i++ {
		num, err := parseDigit(line[i])
		if err != nil {
			fmt.Printf("Error parsing digit: %v\n", err)
			os.Exit(1)
		}
		if num > secondMaxDigit {
			secondMaxDigit = num
		}
	}

	return (firstMaxDigit * 10) + secondMaxDigit
}

func Day3Part1() {
	file, err := os.Open(FILE_PATH)
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

		maxPair := maxPairD1(line)
		pairSum += maxPair
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", pairSum)
}
