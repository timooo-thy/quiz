package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const TIME_LIMIT = 30 * time.Second
const IS_RANDOM = true

func main() {
	file, err :=  os.Open("problems.csv")
	check(err)
	defer file.Close()

	problems := []Problem{}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		problems = append(problems, Problem{
			question: record[0],
			answer: record[1],
		})
	}

	if IS_RANDOM {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	correct := 0
	total := len(problems)

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("Press Enter to start the quiz.")
	fmt.Scanln()
	fmt.Printf("Your %v starts now.\n\n", TIME_LIMIT)
	time.AfterFunc(TIME_LIMIT, func() {
		fmt.Println("\n⏰ Time's up!")
		fmt.Printf("That's the end of the quiz. Your score is %v/%v.", correct, total)
		os.Exit(0)
	})

	for p := range problems {
		var answer string
		fmt.Println(problems[p].question)
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)
		
		if answer == problems[p].answer {
			correct += 1
			fmt.Println("You're right!")
		} else {
			fmt.Println("You're incorrect...")
		}
	}

	fmt.Printf("That's the end of the quiz. Your score is %v/%v.", correct, total)

}