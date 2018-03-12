package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	quizFilename := flag.String("qf", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*quizFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file '%s'.\n", *quizFilename))
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	quiz := make([]QuizItem, len(records))
	for i, v := range records {
		quiz[i] = QuizItem{Question: v[0], Answer: v[1]}
	}

	var ans string
	var correctAns int
	for _, v := range quiz {
		fmt.Printf("%s: ", v.Question)
		fmt.Scanln(&ans)

		if ans == v.Answer {
			correctAns++
		}
	}

	fmt.Printf("Correct answers: %v/%v\n", correctAns, len(quiz))
}

type QuizItem struct {
	Question, Answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}