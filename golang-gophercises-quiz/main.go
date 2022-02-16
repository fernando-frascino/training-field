package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilaname := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")

	flag.Parse()

	file, err := os.Open(*csvFilaname)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilaname))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
