package main

import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file name in format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s", *csvFilename), 1)
	}

	r := csv.NewReader(file)
	questions, _ := r.ReadAll()

	fmt.Println(questions)
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}