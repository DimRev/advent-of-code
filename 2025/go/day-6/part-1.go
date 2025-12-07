package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-6/part-1/input.txt"
)

func mapTransformNums(slice []string) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, num)
	}
	return result
}

func mapCleanSymbols(slice []string) []string {
	result := make([]string, 0)
	for _, v := range slice {
		if v == "" {
			continue
		}
		result = append(result, v)
	}
	return result
}

func calculateTotal(numsList [][]int, symbols []string, lineLen int) int {
	total := 0
	for i := range lineLen {
		currNums := make([]int, 0)
		for _, nums := range numsList {
			currNums = append(currNums, nums[i])
		}
		switch symbols[i] {
		case "+":
			total += addition(currNums)
		case "*":
			total += multiplication(currNums)
		default:
			fmt.Printf("Unknown symbol: %s\n", symbols[i])
			os.Exit(1)
		}
	}

	return total
}

func Day6Part1() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numsList := make([][]int, 0)
	symbols := make([]string, 0)
	lineLen := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		parsed := strings.Split(line, " ")
		transformed := mapTransformNums(parsed)
		if len(transformed) != 0 {
			numsList = append(numsList, transformed)
		} else {
			cleanedSymbols := mapCleanSymbols(parsed)
			symbols = cleanedSymbols
			lineLen = len(cleanedSymbols)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := calculateTotal(numsList, symbols, lineLen)

	fmt.Printf("Solution: %d\n", solution)

}
