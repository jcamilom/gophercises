package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type QuizItem struct {
	Question, Answer string
}

func main() {

	quizFilename := flag.String("qf", "problems.csv", "The name of the quiz file")
	flag.Parse()

	c, err := ioutil.ReadFile(*quizFilename)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(c)))

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	quiz := make([]QuizItem, len(records))
	for i, v := range records {
		quiz[i].Question = v[0]
		quiz[i].Answer = v[1]
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