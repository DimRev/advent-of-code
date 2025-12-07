package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH_P2 = "../inputs/day-6/part-1/input.txt"
)

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func padLines(lines []string) []string {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	padded := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < maxLen {
			padded[i] = line + strings.Repeat(" ", maxLen-len(line))
		} else {
			padded[i] = line
		}
	}
	return padded
}

func isColumnEmpty(lines []string, col int) bool {
	for row := 0; row < len(lines); row++ {
		if col < len(lines[row]) && lines[row][col] != ' ' {
			return false
		}
	}
	return true
}

func readNumberVertically(lines []string, col, startRow int) (int, int, error) {
	numStr := ""
	row := startRow
	for row < len(lines) {
		if col >= len(lines[row]) {
			break
		}
		ch := lines[row][col]
		if ch >= '0' && ch <= '9' {
			numStr += string(ch)
			row++
		} else {
			break
		}
	}
	if numStr == "" {
		return 0, startRow, fmt.Errorf("no number found")
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, startRow, err
	}
	return num, row - 1, nil
}

func processProblem(lines []string, startCol int) ([]int, string, int) {
	var numbers []int
	var operator string
	col := startCol

	for col >= 0 {
		if isColumnEmpty(lines, col) {
			break
		}

		for row := 0; row < len(lines); row++ {
			if col >= len(lines[row]) {
				continue
			}

			char := lines[row][col]
			if char == ' ' {
				continue
			}

			if char == '+' || char == '*' {
				operator = string(char)
			} else if char >= '0' && char <= '9' {
				num, lastRow, err := readNumberVertically(lines, col, row)
				if err == nil {
					numbers = append(numbers, num)
					row = lastRow
				}
			}
		}
		col--
	}
	return numbers, operator, col
}

func calculateProblem(numbers []int, operator string) int {
	if len(numbers) == 0 || operator == "" {
		return 0
	}

	switch operator {
	case "+":
		return addition(numbers)
	case "*":
		return multiplication(numbers)
	default:
		return 0
	}
}

func Day6Part2() {
	file, err := os.Open(FILE_PATH_P2)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(lines) == 0 {
		fmt.Println("No lines found")
		return
	}

	lines = padLines(lines)
	grandTotal := 0
	col := len(lines[0]) - 1

	for col >= 0 {
		for col >= 0 && isColumnEmpty(lines, col) {
			col--
		}

		if col < 0 {
			break
		}

		numbers, operator, nextCol := processProblem(lines, col)
		result := calculateProblem(numbers, operator)
		grandTotal += result
		col = nextCol
	}

	fmt.Printf("Solution: %d\n", grandTotal)
}
