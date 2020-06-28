package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file name in format of question,answer")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("correct answers: %d out of %d\n", correct, len(problems))
}

func parseLines(problems [][]string) []Problem {
	ret := make([]Problem, len(problems))

	for i, problem := range problems {
		ret[i] = Problem{q: problem[0], a: strings.TrimSpace(problem[1])}
	}

	return ret
}

// Problem type
type Problem struct {
	q string
	a string
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
