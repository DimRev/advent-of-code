package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH_P2 = "../inputs/day-5/part-1/input.txt"
)

func formatRangeP2(rangeStr string) (startRange, endRange int, err error) {
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

func sortedInsertRangeP2(ranges []string, rangeToInsert string) ([]string, error) {
	if len(ranges) == 0 {
		return []string{rangeToInsert}, nil
	}

	_, _, err := formatRangeP2(rangeToInsert)
	if err != nil {
		return []string{}, err
	}
	insertStart, _, _ := formatRangeP2(rangeToInsert)

	low := 0
	high := len(ranges) - 1
	insertIdx := len(ranges)

	for low <= high {
		mid := low + (high-low)/2
		midStart, _, err := formatRangeP2(ranges[mid])
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

func squashRangesP2(ranges []string) ([]string, error) {
	if len(ranges) == 0 {
		return ranges, nil
	}

	type rangeData struct {
		start int
		end   int
	}

	parsedRanges := make([]rangeData, 0, len(ranges))
	for _, rangeStr := range ranges {
		start, end, err := formatRangeP2(rangeStr)
		if err != nil {
			return nil, err
		}
		parsedRanges = append(parsedRanges, rangeData{start: start, end: end})
	}

	result := make([]rangeData, 0)
	for _, r := range parsedRanges {
		if len(result) == 0 {
			result = append(result, r)
			continue
		}

		lastIdx := len(result) - 1
		if r.start <= result[lastIdx].end+1 {
			if r.end > result[lastIdx].end {
				result[lastIdx].end = r.end
			}
		} else {
			result = append(result, r)
		}
	}

	squashed := make([]string, 0, len(result))
	for _, r := range result {
		squashed = append(squashed, fmt.Sprintf("%d-%d", r.start, r.end))
	}

	return squashed, nil
}

func checkTotalIdsInRanges(ranges []string) int {
	totalCount := 0

	for _, rangeStr := range ranges {
		startRange, endRange, err := formatRangeP2(rangeStr)
		if err != nil {
			continue
		}
		count := endRange - startRange + 1
		totalCount += count
	}

	return totalCount
}

func Day5Part2() {
	file, err := os.Open(FILE_PATH_P2)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rangesList := make([]string, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		rangesList, err = sortedInsertRangeP2(rangesList, line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	squashed, err := squashRangesP2(rangesList)
	if err != nil {
		fmt.Printf("Error squashing ranges: %v\n", err)
		os.Exit(1)
	}
	freshCount := checkTotalIdsInRanges(squashed)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", freshCount)
}
