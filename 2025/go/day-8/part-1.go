package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func connectPairs(segments []*Segment, numOfConnections int) {
	for i := range numOfConnections {
		segment := segments[i]
		segment.Start.connect(segment.End)
	}
}

func sortIslandsBySize(islandSizes []int) []int {
	sort.Slice(islandSizes, func(i, j int) bool {
		return islandSizes[i] > islandSizes[j]
	})

	return islandSizes
}

func productLargestIslands(islandSizes []int, num int) int {
	product := 0

	for i := range num {
		if i == 0 {
			product = 1
		}
		size := islandSizes[i]
		product *= size
	}
	return product
}

func Day8Part1() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	runningId := 0
	points := make([]*Point, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		point := parseLineIntoPoint(line, runningId)
		points = append(points, point)
		runningId++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	segments := sortSegmentsByLength(parsePointsIntoSegments(points))
	connectPairs(segments, 1000)
	islands := sortIslandsBySize(islandSizes(points))
	solution := productLargestIslands(islands, 3)

	fmt.Printf("Solution: %d\n", solution)

}
