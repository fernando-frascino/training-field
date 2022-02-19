package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	problems := readLines(lines)
	correctCounter := 0

	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, problem.quetion)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correctCounter++
		}
	}

	fmt.Printf("You scored %d out of %d\n", correctCounter, len(problems))
}

func readLines(lines [][]string) []problem {
	result := make([]problem, len(lines))

	for i, line := range lines {
		result[i] = problem{
			quetion: line[0],
			answer:  strings.TrimSpace(line[1]),
		}
	}

	return result
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	quetion string
	answer  string
}
