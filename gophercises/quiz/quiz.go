package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_QUIZ_TIME          = 30
	DEFAULT_PROBLEMS_FILE_NAME = "problems.csv"
)

type Problem struct {
	q string
	a string
}

func parseProblemsFile(filename string) []Problem {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	problems := make([]Problem, 0, len(records))
	for _, v := range records {
		problems = append(problems, Problem{v[0], v[1]})
	}

	return problems
}

func parseFlags() (string, int) {
	var filename string = DEFAULT_PROBLEMS_FILE_NAME
	var time int = DEFAULT_QUIZ_TIME
	for _, v := range os.Args {
		if !strings.Contains(v, "--") {
			continue
		}
		values := strings.Split(v, "=")
		if len(values) != 2 {
			log.Fatal("wrong value for flag")
		}

		switch values[0] {
		case "--filename":
			filename = values[1]
		case "--time":
			t, err := strconv.Atoi(values[1])
			if err != nil {
				log.Fatal("time must be a number")
			}
			time = t
		default:
			log.Fatal("unsupported flag")
		}
	}
	return filename, time
}

func main() {
	filename, t := parseFlags()
	problems := parseProblemsFile(filename)

	totalQuestions := len(problems)
	var correctAnswers int
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %d seconds to answer %d.\nPress Enter to start the timer", t, totalQuestions)
	reader.ReadString('\n')

	timer := time.After(time.Duration(t) * time.Second)

	printResults := func() {
		fmt.Fprintf(os.Stdout, "Total Questions: %d\nCorrect Answers: %d\nWrong Answers: %d\nPercentage: %.2f\n", totalQuestions, correctAnswers, totalQuestions-correctAnswers, float32(correctAnswers)/float32(totalQuestions)*100.0)
	}

	go func() {
		<-timer
		fmt.Println("Time out!")
		printResults()
		os.Exit(0)
	}()

	for _, q := range problems {
		fmt.Println(q.q)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(answer) == q.a {
			correctAnswers += 1
		}
	}
	printResults()
}
