package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		if atEOF {
			return 0, nil, nil
		}
		return 0, nil, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[0:i], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func parseLine(line string) (minNum int, maxNum int, err error) {
	nums := strings.Split(line, "-")
	if len(nums) != 2 {
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

func checkIfValid(num int) bool {
	numStr := strconv.Itoa(num)
	freqMap := make(map[string]int)

	for i := 0; i < len(numStr); i++ {
		digit := string(numStr[i])
		freqMap[digit]++
	}

	allCountsAreTwo := true
	for d, c := range freqMap {
		if c != 2 {
			allCountsAreTwo = false
		}
		fmt.Printf("%s: %d\n", d, c)
	}
	fmt.Println()

	return !allCountsAreTwo
}

func findInvalidNumsInRange(minNum int, maxNum int) (invalidNums []int) {
	invalidNums = make([]int, 0)

	for i := minNum; i <= maxNum; i++ {
		if !checkIfValid(i) {
			invalidNums = append(invalidNums, i)
		}
	}

	return invalidNums
}

func Day2Part1() {
	{
		file, err := os.Open("../inputs/day-2/part-1/example.txt")
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(scanCommaSeparated)

		for scanner.Scan() {
			numRange := strings.TrimSpace(scanner.Text())
			if len(numRange) == 0 {
				continue
			}
			minNum, maxNum, err := parseLine(numRange)
			if err != nil {
				fmt.Printf("Error parsing line: %v\n", err)
				os.Exit(1)
			}
			invalidNums := findInvalidNumsInRange(minNum, maxNum)

			fmt.Printf("min: %d, max: %d, invalid: %v\n", minNum, maxNum, invalidNums)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

	}
}
