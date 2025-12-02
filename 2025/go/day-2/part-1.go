package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanCommaSeparatedP1(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		if atEOF {
			return 0, nil, nil
		}
		return 0, nil, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i] == byte(COMMA_SEPARATOR) {
			return i + 1, data[0:i], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func parseLineP1(line string) (minNum int, maxNum int, err error) {
	nums := strings.Split(line, string(RANGE_SEPARATOR))
	if len(nums) != EXPECTED_RANGE_PARTS {
		return 0, 0, fmt.Errorf("invalid line: %s", line)
	}

	minNum, err = strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, err
	}

	maxNum, err = strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, err
	}

	return minNum, maxNum, nil
}

func checkIfInvalidP1(num int) bool {
	numStr := strconv.Itoa(num)
	length := len(numStr)

	if length%2 != 0 {
		return false
	}

	if numStr[0] == byte(LEADING_ZERO_CHAR) {
		return false
	}

	half := length / 2
	firstHalf := numStr[0:half]
	secondHalf := numStr[half:]

	return firstHalf == secondHalf
}

func findInvalidNumsInRangeP1(minNum int, maxNum int) (invalidNums []int) {
	invalidNums = make([]int, 0)

	for i := minNum; i <= maxNum; i++ {
		if checkIfInvalidP1(i) {
			invalidNums = append(invalidNums, i)
		}
	}

	return invalidNums
}

func sumInvalidNumsListP1(invalidNums []int) int {
	sum := 0
	for _, num := range invalidNums {
		sum += num
	}
	return sum
}

func Day2Part1() {
	{
		file, err := os.Open("../inputs/day-2/part-1/input.txt")
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(scanCommaSeparatedP1)

		invalidNums := make([]int, 0)
		for scanner.Scan() {
			numRange := strings.TrimSpace(scanner.Text())
			if len(numRange) == 0 {
				continue
			}
			minNum, maxNum, err := parseLineP1(numRange)
			if err != nil {
				fmt.Printf("Error parsing line: %v\n", err)
				os.Exit(1)
			}
			in := findInvalidNumsInRangeP1(minNum, maxNum)
			invalidNums = append(invalidNums, in...)

		}
		solution := sumInvalidNumsListP1(invalidNums)
		fmt.Printf("Solution: %d\n", solution)

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

	}
}
