package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-5/part-1/input.txt"
)

type Range struct {
	Start, End int
}

func parseRange(rangeStr string) (Range, error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return Range{}, fmt.Errorf("invalid range format: %s", rangeStr)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, err
	}

	return Range{Start: start, End: end}, nil
}

func isIDInRanges(ranges []Range, id int) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
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

	var ranges []Range
	var ids []int
	isRangeSection := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isRangeSection = false
			continue
		}
		if isRangeSection {
			r, err := parseRange(line)
			if err != nil {
				fmt.Printf("Error parsing range: %v\n", err)
				os.Exit(1)
			}
			ranges = append(ranges, r)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Error parsing ID: %v\n", err)
				os.Exit(1)
			}
			ids = append(ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	count := 0
	for _, id := range ids {
		if isIDInRanges(ranges, id) {
			count++
		}
	}

	fmt.Printf("Solution: %d\n", count)
}
