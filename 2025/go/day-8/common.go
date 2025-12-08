package day8

import (
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

func islandSizesP2(points []*Point) []int {
	visitedIds := make(map[int]bool)
	islandSizes := make([]int, 0)
	for _, p := range points {
		size := traverseConnections(p, 0, visitedIds)
		if size > 0 {
			islandSizes = append(islandSizes, size)
		}
		if len(islandSizes) > 1 {
			return []int{}
		}
	}
	return islandSizes
}
