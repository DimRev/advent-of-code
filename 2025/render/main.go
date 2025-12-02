package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	fmt.Print("Day/Part    ")
	for _, lang := range languages {
		fmt.Printf("%-20s", lang.name)
	}
	fmt.Println()

	fmt.Print("------------")
	for range languages {
		fmt.Print("--------------------")
	}
	fmt.Println()

	for _, col := range sortedColumns {
		fmt.Printf("%-12s", col)
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
			fmt.Printf("%-20s", formatted)
		}
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
