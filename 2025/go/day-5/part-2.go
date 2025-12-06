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
	FILE_PATH_P2 = "../inputs/day-5/part-1/input.txt"
)

type RangeP2 struct {
	Start, End int
}

func parseRangeP2(rangeStr string) (RangeP2, error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return RangeP2{}, fmt.Errorf("invalid range format: %s", rangeStr)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return RangeP2{}, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return RangeP2{}, err
	}

	return RangeP2{Start: start, End: end}, nil
}

func squashRanges(ranges []RangeP2) []RangeP2 {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	result := []RangeP2{ranges[0]}

	for _, r := range ranges[1:] {
		last := &result[len(result)-1]
		if r.Start <= last.End+1 {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			result = append(result, r)
		}
	}

	return result
}

func countTotalIDs(ranges []RangeP2) int {
	total := 0
	for _, r := range ranges {
		total += r.End - r.Start + 1
	}
	return total
}

func Day5Part2() {
	file, err := os.Open(FILE_PATH_P2)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var ranges []RangeP2
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		r, err := parseRangeP2(line)
		if err != nil {
			fmt.Printf("Error parsing range: %v\n", err)
			os.Exit(1)
		}
		ranges = append(ranges, r)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	squashed := squashRanges(ranges)
	totalCount := countTotalIDs(squashed)

	fmt.Printf("Solution: %d\n", totalCount)
}
