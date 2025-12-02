package main

import (
	"fmt"
	"os"
	"time"

	day1 "github.com/DimRev/advent-of-code/day-1"
	day2 "github.com/DimRev/advent-of-code/day-2"
	day3 "github.com/DimRev/advent-of-code/day-3"
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
)

var commands = []Command{D1P1, D1P2}

var cmdMap = map[Command]func(){
	D1P1: day1.Day1Part1,
	D1P2: day1.Day1Part2,

	D2P1: day2.Day2Part1,
	D2P2: day2.Day2Part2,

	D3P1: day3.Day3Part1,
	D3P2: day3.Day3Part2,
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
