package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-9/part-1/input.txt"
)

type Point struct {
	Id int
	X  int
	Y  int
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

type Polygon struct {
	Points []*Point
}

func (p *Polygon) String() string {
	str := "Polygon("
	for i, point := range p.Points {
		str += point.String()
		if i < len(p.Points)-1 {
			str += ", "
		}
	}
	return str + ")"
}

func (p *Polygon) CheckIfPointInside(po *Point) bool {
	n := len(p.Points)
	inside := false

	// Handle edge case - if point is on a vertex, consider it outside
	for i := range n {
		if p.Points[i].X == po.X && p.Points[i].Y == po.Y {
			return false
		}
	}

	// Ray casting algorithm
	for i := range n {
		p1 := p.Points[i]
		p2 := p.Points[(i+1)%n]

		// Ensure p1 has lower y-coordinate than p2
		if p1.Y > p2.Y {
			p1, p2 = p2, p1
		}

		// Check if the horizontal ray from point intersects the edge
		if po.Y <= p2.Y {
			if po.Y > p1.Y {
				// Calculate x-coordinate of intersection
				if po.X <= p1.X+(p2.X-p1.X)*(po.Y-p1.Y)/(p2.Y-p1.Y) {
					inside = !inside
				}
			}
		}
	}

	return inside
}

type Rectangle struct {
	TopLeft     *Point
	TopRight    *Point
	BottomRight *Point
	BottomLeft  *Point

	MinX int
	MaxX int
	MinY int
	MaxY int

	Area int
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%s, %s, %s, %s)", r.TopLeft, r.TopRight, r.BottomRight, r.BottomLeft)
}

func GenerateRectangleFrom2Points(p1, p2 *Point) *Rectangle {
	if p1.X == p2.X || p1.Y == p2.Y {
		// Degenerate case - points are aligned horizontally or vertically
		return nil // or handle as needed
	}

	// Calculate min/max coordinates
	minX := int(math.Min(float64(p1.X), float64(p2.X)))
	maxX := int(math.Max(float64(p1.X), float64(p2.X)))
	minY := int(math.Min(float64(p1.Y), float64(p2.Y)))
	maxY := int(math.Max(float64(p1.Y), float64(p2.Y)))

	// Define corners correctly
	topLeft := &Point{X: minX, Y: maxY}
	topRight := &Point{X: maxX, Y: maxY}
	bottomLeft := &Point{X: minX, Y: minY}
	bottomRight := &Point{X: maxX, Y: minY}

	return &Rectangle{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
		MinX:        minX, MinY: minY,
		MaxX: maxX, MaxY: maxY,
		Area: CalculateRectArea(bottomLeft, topRight),
	}
}

func ParseLineIntoPoint(line string, id int) *Point {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		fmt.Printf("Error parsing line: %v\n", line)
		os.Exit(1)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Error parsing line: %v\n", line)
		os.Exit(1)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Error parsing line: %v\n", line)
		os.Exit(1)
	}
	return &Point{X: x, Y: y, Id: id}
}

func CalculateRectArea(p1, p2 *Point) int {
	totalX := math.Abs(float64(p1.X-p2.X)) + 1
	totalY := math.Abs(float64(p1.Y-p2.Y)) + 1
	return int(totalX * totalY)
}

// ReadPointsFromFile reads points from the input file
func ReadPointsFromFile(filePath string) ([]*Point, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	points := make([]*Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points[i] = &Point{X: x, Y: y, Id: i}
	}

	return points, nil
}

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sort returns two integers in ascending order
func Sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// ManhattanDistance calculates the Manhattan distance between two points
func ManhattanDistance(a, b *Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

// RectangleArea calculates the area of a rectangle from two corner points
func RectangleArea(x1, y1, x2, y2 int) int {
	width := Abs(x2-x1) + 1
	height := Abs(y2-y1) + 1
	return width * height
}
