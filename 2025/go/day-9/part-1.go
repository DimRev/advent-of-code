package day9

import (
	"fmt"
	"os"
)

func calculateMaxArea(points []*Point) int {
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area := CalculateRectArea(p1, p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func Day9Part1() {
	points, err := ReadPointsFromFile(FILE_PATH)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := calculateMaxArea(points)

	fmt.Printf("Solution: %d\n", solution)
}
