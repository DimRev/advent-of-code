package main

import (
	"fmt"
	"os"

	day1 "github.com/DimRev/advent-of-code/day-1"
)

type Command string

const (
	D1P1 Command = "d1p1"
	D1P2 Command = "d1p2"
)

var commands = []Command{D1P1, D1P2}

var cmdMap = map[Command]func(){
	D1P1: day1.Day1Part1,
	D1P2: day1.Day1Part2,
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

	cmdMap[cmd]()
}
