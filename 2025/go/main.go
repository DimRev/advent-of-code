package main

import (
	"fmt"
	"os"
	"time"

	day1 "github.com/DimRev/advent-of-code/day-1"
	day2 "github.com/DimRev/advent-of-code/day-2"
	day3 "github.com/DimRev/advent-of-code/day-3"
	day4 "github.com/DimRev/advent-of-code/day-4"
	day5 "github.com/DimRev/advent-of-code/day-5"
	day6 "github.com/DimRev/advent-of-code/day-6"
	day7 "github.com/DimRev/advent-of-code/day-7"
	day8 "github.com/DimRev/advent-of-code/day-8"
	day9 "github.com/DimRev/advent-of-code/day-9"
	"github.com/DimRev/advent-of-code/lib"
)

type Command string

const (
	D1P1 Command = "d1p1"
	D1P2 Command = "d1p2"

	D2P1 Command = "d2p1"
	D2P2 Command = "d2p2"

	D3P1 Command = "d3p1"
	D3P2 Command = "d3p2"

	D4P1 Command = "d4p1"
	D4P2 Command = "d4p2"

	D5P1 Command = "d5p1"
	D5P2 Command = "d5p2"

	D6P1 Command = "d6p1"
	D6P2 Command = "d6p2"

	D7P1 Command = "d7p1"
	D7P2 Command = "d7p2"

	D8P1 Command = "d8p1"
	D8P2 Command = "d8p2"

	D9P1 Command = "d9p1"
	D9P2 Command = "d9p2"
)

var commands = []Command{D1P1, D1P2}

var cmdMap = map[Command]func(){
	D1P1: day1.Day1Part1,
	D1P2: day1.Day1Part2,

	D2P1: day2.Day2Part1,
	D2P2: day2.Day2Part2,

	D3P1: day3.Day3Part1,
	D3P2: day3.Day3Part2,

	D4P1: day4.Day4Part1,
	D4P2: day4.Day4Part2,

	D5P1: day5.Day5Part1,
	D5P2: day5.Day5Part2,

	D6P1: day6.Day6Part1,
	D6P2: day6.Day6Part2,

	D7P1: day7.Day7Part1,
	D7P2: day7.Day7Part2,

	D8P1: day8.Day8Part1,
	D8P2: day8.Day8Part2,

	D9P1: day9.Day9Part1,
	D9P2: day9.Day9Part2,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		fmt.Println("Available commands:")
		for _, cmd := range commands {
			fmt.Printf("\t- %s\n", cmd)
		}
		os.Exit(1)
	}

	cmdStr := os.Args[1]
	cmd := Command(cmdStr)

	if _, ok := cmdMap[cmd]; !ok {
		fmt.Printf("Unknown command: %s\n", cmdStr)
		fmt.Println("Available commands:")
		for _, cmdEnum := range commands {
			fmt.Printf("\t- %s\n", cmdEnum)
		}
		os.Exit(1)
	}

	startTs := time.Now()
	cmdMap[cmd]()
	endTs := time.Now()
	elapsed := endTs.Sub(startTs)
	fmt.Printf("Finished running %s in %d(μs)\n", cmdStr, elapsed.Microseconds())

	err := lib.PopulateRenderer(cmdStr, int(elapsed.Microseconds()))
	if err != nil {
		fmt.Printf("Error populating renderer: %v\n", err)
	}
}
