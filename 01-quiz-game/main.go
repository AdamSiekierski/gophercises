package main

import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file name in format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s", *csvFilename), 1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("failed to parse the csv file: %s", *csvFilename), 1)
	}

	problems := parseLines(lines)

	correct := 0

	for i, p := range problems {
		fmt.Printf("problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("correct answers: %d out of %d\n", correct, len(problems))
}

func parseLines(problems [][]string) []Problem {
	ret := make([]Problem, len(problems))

	for i, problem := range problems {
		ret[i] = Problem{q: problem[0], a: problem[1]}
	}

	return ret
}

type Problem struct {
	q string
	a string
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}