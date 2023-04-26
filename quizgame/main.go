package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	_ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse provided CSV file: %s", *csvFileName))
	}

	problems := parseLines(lines)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Println("Correct")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type Problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
