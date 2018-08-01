package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var fileFlag string
var timerFlag int
var shuffleFlag bool

// Problem contains a question and an answer
type Problem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func init() {
	flag.StringVar(&fileFlag, "f", "problems.csv", "The CSV file you wish to use for questions and answers")
	flag.IntVar(&timerFlag, "t", 30, "Seconds given to finish the quiz.")
	flag.BoolVar(&shuffleFlag, "s", false, "Shuffles the questions")
}

func main() {
	flag.Parse()
	csvFile, _ := os.Open(fileFlag)
	csvRows := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem

	for {
		line, error := csvRows.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}

	if shuffleFlag == true {
		for p := range problems {
			j := rand.Intn(p + 1)
			problems[p], problems[j] = problems[j], problems[p]
		}
	}

	score := 0
	totalQuestions := len(problems)
	quit := make(chan struct{})
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %v seconds to complete the quiz. Press enter to begin. ", timerFlag)
	reader.ReadString('\n')

	ticker := time.NewTicker(time.Duration(timerFlag) * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("Time has run out. The quiz has ended. Press enter to see your score.")
			close(quit)
		}
	}()

	fmt.Print("Quiz Started \n")

quizLoop:
	for _, prob := range problems {
		select {
		case <-quit:
			break quizLoop
		default:
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("%v \n", prob.Question)
			text, _ := reader.ReadString('\n')
			if strings.TrimRight(text, "\n") == prob.Answer {
				score++
			}
		}
	}

	fmt.Printf("You got %v out of %v right. \n", score, totalQuestions)
}
