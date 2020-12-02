package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read a csv file
	var fileName string
	flag.StringVar(&fileName, "file", "problems.csv", "File to find quiz")
	flag.Parse()
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	quizReader := csv.NewReader(file)
	quizEntries, err := quizReader.ReadAll()
	if err != nil {
		panic(err)
	}
	inputReader := bufio.NewReader(os.Stdin)
	// looping through all the entries
	var right int
	var wrong int

	for _, entry := range quizEntries {
		if len(entry) != 2 {
			panic("Yo question is bad")
		}
		question := entry[0]
		answer := entry[1]
		// Output the question
		fmt.Println(question)
		// read a user input
		userText, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		// Compare it to "answer"
		if answer == strings.TrimSpace(userText) {
			right++
			fmt.Println("You got it right!")
		} else {
			wrong++
			fmt.Println("That's wrong!")
		}
		// Keep track of right/wrong
	}
	// Ending loop: output score
	fmt.Printf("Finished! You got %d right and %d wrong!\n", right, wrong)
}
