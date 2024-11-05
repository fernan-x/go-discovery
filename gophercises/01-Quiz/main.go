package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Question struct {
	question string;
	answer string;
}

func readQuiz(filename string) []Question {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open file", err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	var questions []Question
	for _, record := range records {
		question := record[0]
		answer := record[1]
		questions = append(questions, Question{question, answer})
	}

	return questions
}

func askQuestion(question string) string {
	var answer string
	fmt.Println("What is", question, "?")
	reader := bufio.NewReader(os.Stdin)
	answer, _ = reader.ReadString('\n')
	return strings.Trim(answer, "\n")
}

func main() {
	var correct = 0;
	questions := readQuiz("data.csv")
	for _, question := range questions {
		response := askQuestion(question.question)

		if response == question.answer {
			correct++
		}
	}

	fmt.Println("Your score is", correct, "/", len(questions))
}