package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8Part2() {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_ = strings.TrimSpace(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	solution := 0

	fmt.Printf("Solution: %d\n", solution)

}
