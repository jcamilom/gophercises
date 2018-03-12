package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	quizFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timerLimit := flag.Int("limit", 30, "Time limit in seconds to answer each question")
	flag.Parse()

	if *timerLimit < 1 {
		exit("Failed to run the quiz. Time limit should be bigger than 0s.")
	}

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

	var correctAns int
	var ans string
	fmt.Print("Press enter when you are ready to start the quiz...")
	fmt.Scanln(&ans)
	timer := time.NewTimer(time.Duration(*timerLimit) * time.Second)

	problemloop:
	for _, v := range quiz {
		fmt.Printf("%s: ", v.Question)
		answerCh := make(chan string)
		go func() {
			fmt.Scanln(&ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Print("\nTime expired. ")
			break problemloop
		case ans := <-answerCh:
			if ans == v.Answer {
				correctAns++
			}
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