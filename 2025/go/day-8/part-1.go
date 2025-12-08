package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-8/part-1/input.txt"
)

type ConnectionsMap map[int]*Point

type Point struct {
	Id int
	X  int
	Y  int
	Z  int
	ConnectionsMap
}

func (p *Point) String() string {
	connections := p.listConnectionsIds()
	return fmt.Sprintf("Point[%d](%d, %d, %d) Connections: %v", p.Id, p.X, p.Y, p.Z, connections)
}

func (p *Point) distance(target *Point) int {
	diffXSquared := math.Pow(float64(target.X-p.X), 2)
	diffYSquared := math.Pow(float64(target.Y-p.Y), 2)
	diffZSquared := math.Pow(float64(target.Z-p.Z), 2)
	return int(math.Abs(math.Sqrt(diffXSquared + diffYSquared + diffZSquared)))
}

func (p *Point) connect(target *Point) {
	_, ok := p.ConnectionsMap[target.Id]
	if !ok {
		p.ConnectionsMap[target.Id] = target
	}

	_, ok = target.ConnectionsMap[p.Id]
	if !ok {
		target.ConnectionsMap[p.Id] = p
	}
}

func (p *Point) listConnectionsIds() []int {
	ids := make([]int, 0)
	for id := range p.ConnectionsMap {
		ids = append(ids, id)
	}
	return ids
}

type Segment struct {
	Start  *Point
	End    *Point
	Length int
}

func (s *Segment) String() string {
	return fmt.Sprintf("Segment(Point[%d], Point[%d], %d)", s.Start.Id, s.End.Id, s.Length)
}

func generatePoint(x, y, z, id int) *Point {
	connectionMap := make(ConnectionsMap)
	return &Point{Id: id, X: x, Y: y, Z: z, ConnectionsMap: connectionMap}
}

func parseLineIntoPoint(line string, id int) *Point {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
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
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Error parsing line: %v\n", line)
		os.Exit(1)
	}
	return generatePoint(x, y, z, id)
}

func parsePointsIntoSegments(points []*Point) []*Segment {
	segments := make([]*Segment, 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			segment := &Segment{
				Start:  points[i],
				End:    points[j],
				Length: points[j].distance(points[i]),
			}

			segments = append(segments, segment)
		}
	}

	return segments
}

func sortSegmentsByLength(segments []*Segment) []*Segment {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Length < segments[j].Length
	})

	return segments
}

func connectPairs(segments []*Segment, numOfConnections int) {
	for i := 0; i < numOfConnections; i++ {
		segment := segments[i]
		segment.Start.connect(segment.End)
	}
}

func islandSizes(points []*Point) []int {
	visitedIds := make(map[int]bool)
	islandSizes := make([]int, 0)
	for _, p := range points {
		size := traverseConnections(p, 0, visitedIds)
		if size > 0 {
			islandSizes = append(islandSizes, size)
		}
	}

	return islandSizes
}

func traverseConnections(point *Point, prevCount int, visitedIds map[int]bool) int {
	_, ok := visitedIds[point.Id]
	if ok {
		return 0
	}

	if len(point.ConnectionsMap) == 0 {
		return 1
	}

	count := 1
	visitedIds[point.Id] = true

	for _, p := range point.ConnectionsMap {
		count += traverseConnections(p, prevCount+1, visitedIds)
	}

	return count
}

func sortIslandsBySize(islandSizes []int) []int {
	sort.Slice(islandSizes, func(i, j int) bool {
		return islandSizes[i] > islandSizes[j]
	})

	return islandSizes
}

func productLargestIslands(islandSizes []int, num int) int {
	product := 0

	for i := 0; i < num; i++ {
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
	segments := make([]*Segment, 0)

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

	segments = sortSegmentsByLength(parsePointsIntoSegments(points))
	connectPairs(segments, 1000)
	islands := sortIslandsBySize(islandSizes(points))
	solution := productLargestIslands(islands, 3)

	fmt.Printf("Solution: %d\n", solution)

}
