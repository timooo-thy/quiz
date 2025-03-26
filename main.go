package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
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

func waitForEnter() {
    fmt.Println("Press Enter to start the quiz")
    reader := bufio.NewReader(os.Stdin)
    _, _ = reader.ReadString('\n')
}

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

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("Press Enter to start the quiz")
	fmt.Scanln()

	correct := 0
	total := len(problems)

	for p := range problems {
		var answer string
		fmt.Println(problems[p].question)
		fmt.Scanln(&answer)
		if answer == problems[p].answer {
			correct += 1
			fmt.Println("You're right!")
		} else {
			fmt.Println("You're incorrect...")
		}
	}

	fmt.Printf("That's the end of the quiz. Your score is %v/%v.", correct, total)

}