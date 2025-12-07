package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE_PATH = "../inputs/day-7/part-1/input.txt"
)

func generateBeamList(line string) []bool {
	result := make([]bool, len(line))
	for i, r := range line {
		if r == 'S' {
			result[i] = true
			continue
		}
		result[i] = false
	}
	return result
}

func progressSplitters(line string, beamList []bool) (newBeamList []bool, splitCount int) {
	newBeamList = make([]bool, len(line))
	splitCount = 0
	for i, r := range line {
		if r == '^' && beamList[i] {
			if i-1 >= 0 {
				newBeamList[i-1] = true
			}
			if i+1 < len(line) {
				newBeamList[i+1] = true
			}
			newBeamList[i] = false
			splitCount++
			continue
		}
		newBeamList[i] = newBeamList[i] || beamList[i]
	}
	return newBeamList, splitCount
}

func formatBeamList(beamList []bool) string {
	result := ""
	for _, b := range beamList {
		if b {
			result += "|"
		} else {
			result += "."
		}
	}
	return result
}

func Day7Part1() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	shouldPrint := false

	splitCount := 0
	beamList := make([]bool, 0)
	idx := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if idx == 0 {
			beamList = generateBeamList(line)
			if shouldPrint {
				fmt.Println(line)
				fmt.Println(formatBeamList(beamList))
			}
		} else if idx%2 == 0 {
			newBeamList, newSplitCount := progressSplitters(line, beamList)
			splitCount += newSplitCount
			beamList = newBeamList
			if shouldPrint {
				fmt.Println(line)
				fmt.Println(formatBeamList(beamList))
			}
		}
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := splitCount

	fmt.Printf("Solution: %d\n", solution)

}
