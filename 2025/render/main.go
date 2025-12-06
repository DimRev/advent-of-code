package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Data struct {
	Go     map[string]interface{} `json:"go"`
	Python map[string]interface{} `json:"python"`
	Ts     map[string]interface{} `json:"ts"`
}

type Command string

const (
	Init     Command = "init"
	Render   Command = "render"
	Populate Command = "populate"
)

var dayParts = []string{
	"d1p1", "d1p2",
	"d2p1", "d2p2",
	"d3p1", "d3p2",
	"d4p1", "d4p2",
	"d5p1", "d5p2",
	"d6p1", "d6p2",
}

func initData() error {
	data := Data{
		Go:     make(map[string]interface{}),
		Python: make(map[string]interface{}),
		Ts:     make(map[string]interface{}),
	}

	for _, dayPart := range dayParts {
		data.Go[dayPart] = "N/A"
		data.Python[dayPart] = "N/A"
		data.Ts[dayPart] = "N/A"
	}

	err := os.MkdirAll("./data", 0755)
	if err != nil {
		return fmt.Errorf("error creating data directory: %v", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile("./data/data.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("Initialized data.json with template")
	return nil
}

func formatNumberWithCommas(num string) string {
	if num == "N/A" {
		return num
	}

	n := len(num)
	if n <= 3 {
		return num
	}

	var result string
	for i, digit := range num {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}
	return result
}

func padCell(content string, width int) string {
	return fmt.Sprintf(" %-*s|", width-2, content)
}

func padSeparator(width int) string {
	return fmt.Sprintf("%s|", strings.Repeat("-", width-1))
}

func populateData(lang, dayPart string, value int) error {
	err := os.MkdirAll("./data", 0755)
	if err != nil {
		return fmt.Errorf("error creating data directory: %v", err)
	}

	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("data.json not found, initializing...")
			err = initData()
			if err != nil {
				return err
			}
			data, err = os.ReadFile("./data/data.json")
			if err != nil {
				return fmt.Errorf("error reading file: %v", err)
			}
		} else {
			return fmt.Errorf("error reading file: %v", err)
		}
	}

	var parsedData Data
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	switch lang {
	case "go":
		parsedData.Go[dayPart] = value
	case "python":
		parsedData.Python[dayPart] = value
	case "ts":
		parsedData.Ts[dayPart] = value
	default:
		return fmt.Errorf("unknown language: %s (valid: go, python, ts)", lang)
	}

	jsonData, err := json.MarshalIndent(parsedData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile("./data/data.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Updated data.json with %s %s to %d\n", lang, dayPart, value)
	return nil
}

func renderData() error {
	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var parsedData Data
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	columnKeys := make(map[string]bool)
	for key := range parsedData.Go {
		columnKeys[key] = true
	}

	var sortedColumns []string
	for key := range columnKeys {
		sortedColumns = append(sortedColumns, key)
	}
	sort.Strings(sortedColumns)

	languages := []struct {
		name string
		data map[string]interface{}
	}{
		{"Go", parsedData.Go},
		{"Python", parsedData.Python},
		{"TypeScript", parsedData.Ts},
	}

	fmt.Println()
	fmt.Println()
	fmt.Print(padCell("Day/Part", 14))
	for _, lang := range languages {
		fmt.Print(padCell(lang.name, 19))
	}
	fmt.Print(padCell("Analytics", 50))
	fmt.Println()

	fmt.Print(padSeparator(14))
	for range languages {
		fmt.Print(padSeparator(19))
	}
	fmt.Print(padSeparator(50))
	fmt.Println()

	for _, col := range sortedColumns {
		fmt.Print(padCell(col, 14))

		// Collect numeric values for analytics
		var values []float64
		for _, lang := range languages {
			value := lang.data[col]
			var numVal float64
			switch v := value.(type) {
			case float64:
				numVal = v
			case int:
				numVal = float64(v)
			default:
				numVal = -1
			}
			values = append(values, numVal)
		}

		// Find minimum value (fastest)
		minVal := values[0]
		if minVal > 0 {
			for _, v := range values[1:] {
				if v > 0 && v < minVal {
					minVal = v
				}
			}
		}

		// Print values
		for _, lang := range languages {
			value := lang.data[col]
			var formatted string
			switch v := value.(type) {
			case float64:
				formatted = fmt.Sprintf("%d", int64(v))
			case int:
				formatted = fmt.Sprintf("%d", v)
			default:
				formatted = fmt.Sprintf("%v", v)
			}
			formatted = formatNumberWithCommas(formatted)
			fmt.Print(padCell(formatted, 19))
		}

		// Print analytics
		var analytics string
		if minVal > 0 {
			parts := []string{}
			for i, lang := range languages {
				if values[i] > 0 {
					ratio := values[i] / minVal
					parts = append(parts, fmt.Sprintf("%s(%.2fx)", lang.name, ratio))
				}
			}
			analytics = ""
			for i, part := range parts {
				if i > 0 {
					analytics += " "
				}
				analytics += fmt.Sprintf("%-12s", part)
			}
		} else {
			analytics = "N/A"
		}
		fmt.Print(padCell(analytics, 50))
		fmt.Println()
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		fmt.Println("Available commands:")
		fmt.Println("\t- init              : Initialize data.json with template")
		fmt.Println("\t- render            : Render the data.json as a table")
		fmt.Println("\t- populate <lang> <dayPart> <value> : Update a specific entry")
		os.Exit(1)
	}

	cmdStr := os.Args[1]
	cmd := Command(cmdStr)

	var err error
	switch cmd {
	case Init:
		err = initData()
	case Render:
		err = renderData()
	case Populate:
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go populate <lang> <dayPart> <value>")
			fmt.Println("Example: go run main.go populate go d1p1 12345")
			os.Exit(1)
		}
		lang := os.Args[2]
		dayPart := os.Args[3]
		var value int
		_, err = fmt.Sscanf(os.Args[4], "%d", &value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: value must be an integer\n")
			os.Exit(1)
		}
		err = populateData(lang, dayPart, value)
	default:
		fmt.Printf("Unknown command: %s\n", cmdStr)
		fmt.Println("Available commands:")
		fmt.Println("\t- init              : Initialize data.json with template")
		fmt.Println("\t- render            : Render the data.json as a table")
		fmt.Println("\t- populate <lang> <dayPart> <value> : Update a specific entry")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
