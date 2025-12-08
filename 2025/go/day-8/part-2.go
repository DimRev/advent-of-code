package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func connectAllPairs(segments []*Segment, points []*Point) int {
	for i := range len(segments) {
		segment := segments[i]
		segment.Start.connect(segment.End)

		islands := islandSizesP2(points)

		if len(islands) == 1 {
			return segment.Start.X * segment.End.X
		}
	}

	return -1
}

func Day8Part2() {
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

	segments := sortSegmentsByLength(parsePointsIntoSegments(points))
	solution := connectAllPairs(segments, points)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solution: %d\n", solution)

}
