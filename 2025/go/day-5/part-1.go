package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-5/part-1/input.txt"
)

func formatRange(rangeStr string) (startRange, endRange int, err error) {
	split := strings.Split(rangeStr, "-")
	if len(split) != 2 {
		err = fmt.Errorf("invalid range format: %s", rangeStr)
		return
	}

	startRange, err = strconv.Atoi(split[0])
	if err != nil {
		return
	}

	endRange, err = strconv.Atoi(split[1])
	if err != nil {
		return
	}
	return startRange, endRange, nil
}

func sortedInsertRange(ranges []string, rangeToInsert string) ([]string, error) {
	if len(ranges) == 0 {
		return []string{rangeToInsert}, nil
	}

	_, _, err := formatRange(rangeToInsert)
	if err != nil {
		return []string{}, err
	}
	insertStart, _, _ := formatRange(rangeToInsert)

	low := 0
	high := len(ranges) - 1
	insertIdx := len(ranges)

	for low <= high {
		mid := low + (high-low)/2
		midStart, _, err := formatRange(ranges[mid])
		if err != nil {
			low = mid + 1
			continue
		}

		if insertStart <= midStart {
			insertIdx = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	result := make([]string, 0, len(ranges)+1)
	result = append(result, ranges[:insertIdx]...)
	result = append(result, rangeToInsert)
	result = append(result, ranges[insertIdx:]...)

	return result, nil
}

func sortedInsertId(ids []string, idToInsert string) ([]string, error) {
	if len(ids) == 0 {
		return []string{idToInsert}, nil
	}

	insertId, err := strconv.Atoi(idToInsert)
	if err != nil {
		return []string{}, err
	}

	low := 0
	high := len(ids) - 1
	insertIdx := len(ids)

	for low <= high {
		mid := low + (high-low)/2

		midId, err := strconv.Atoi(ids[mid])
		if err != nil {
			low = mid + 1
			continue
		}

		if insertId <= midId {
			insertIdx = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	result := make([]string, 0, len(ids)+1)
	result = append(result, ids[:insertIdx]...)
	result = append(result, idToInsert)
	result = append(result, ids[insertIdx:]...)

	return result, nil
}

func findInRanges(ranges []string, idStr string) bool {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return false
	}
	for _, rangeStr := range ranges {
		startRange, endRange, err := formatRange(rangeStr)
		if err != nil {
			continue
		}

		if id < startRange {
			continue
		}
		if id > endRange {
			continue
		}

		if startRange <= id && id <= endRange {
			return true
		}
	}
	return false
}

func Day5Part1() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freshCount := 0
	rangesList := make([]string, 0)
	idsList := make([]string, 0)

	isRange := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isRange = false
			continue
		}
		if isRange {
			rangesList, err = sortedInsertRange(rangesList, line)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		} else {
			idsList, err = sortedInsertId(idsList, line)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}
	}

	for _, idStr := range idsList {
		found := findInRanges(rangesList, idStr)
		if found {
			freshCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", freshCount)
}
