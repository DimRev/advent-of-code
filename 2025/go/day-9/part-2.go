package day9

import (
	"fmt"
	"os"
)

type Edge struct {
	x1, y1, x2, y2 int
}

func buildPolygonEdges(points []*Point) []Edge {
	edges := make([]Edge, 0, len(points))

	for i := 0; i < len(points)-1; i++ {
		edges = append(edges, Edge{
			x1: points[i].X,
			y1: points[i].Y,
			x2: points[i+1].X,
			y2: points[i+1].Y,
		})
	}

	edges = append(edges, Edge{
		x1: points[len(points)-1].X,
		y1: points[len(points)-1].Y,
		x2: points[0].X,
		y2: points[0].Y,
	})

	return edges
}

func rectangleIntersectsEdges(minX, minY, maxX, maxY int, edges []Edge) bool {
	for _, edge := range edges {
		eMinX, eMaxX := Sort(edge.x1, edge.x2)
		eMinY, eMaxY := Sort(edge.y1, edge.y2)

		if minX < eMaxX && maxX > eMinX && minY < eMaxY && maxY > eMinY {
			return true
		}
	}
	return false
}

func calculateMaxBoundedArea(points []*Point) int {
	maxArea := 0
	edges := buildPolygonEdges(points)

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			manhattanDist := ManhattanDistance(p1, p2)
			if manhattanDist*manhattanDist <= maxArea {
				continue
			}

			minX, maxX := Sort(p1.X, p2.X)
			minY, maxY := Sort(p1.Y, p2.Y)

			if !rectangleIntersectsEdges(minX, minY, maxX, maxY, edges) {
				area := RectangleArea(p1.X, p1.Y, p2.X, p2.Y)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func Day9Part2() {
	points, err := ReadPointsFromFile(FILE_PATH)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := calculateMaxBoundedArea(points)

	fmt.Printf("Solution: %d\n", solution)
}
