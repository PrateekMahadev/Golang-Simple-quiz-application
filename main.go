package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// problem struct to hold question, options, and answer
type problem struct {
	q       string   // question
	options []string // options (A, B, C, D)
	a       string   // correct answer
}

// problemPuller reads problems from a CSV file
// This function will return a slice of the problems or an error
func problemPuller(fileName string) ([]problem, error) {
	//read all the problems from the quiz.csv
	// Open the file
	fObj, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening %s file: %s", fileName, err.Error())
	}
	defer fObj.Close()
	// 2. we will create a new reader
	// Create a new CSV reader
	csvR := csv.NewReader(fObj)

	// Read all the lines from CSV
	Clines, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading data in csv format from %s file: %s", fileName, err.Error())
	}

	// Call parseProblems to convert CSV lines to problem structs
	return parseProblems(Clines), nil
}

// it takes lines from the problemPuller as input
// parseProblems converts CSV lines to problem structs
func parseProblems(lines [][]string) []problem {
	// go over the lines and parse the tree with problem struct
	problems := make([]problem, len(lines))
	for i, line := range lines {
		// Ensure we have at least two columns (question and answer)
		if len(line) >= 2 {
			problems[i] = problem{
				q:       line[0],
				a:       strings.TrimSpace(line[1]),
				options: line[2:],
			}
		}
	}
	return problems
}

func main() {
	// Parse command-line flags
	// 1. Input the name of the file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	// 2. Set the duration of the timer
	timer := flag.Int("t", 30, "time for quiz in seconds")
	flag.Parse()

	// Pull problems from CSV file
	// 3. pull the problems from the file (calling our problem puller function)
	problems, err := problemPuller(*fName)
	// 4. Handle the error
	if err != nil {
		exit(fmt.Sprintf("Error: %s", err))
	}

	// Initialize timer
	// 5. Using the duration of the timer, we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)

	// Variables to track correct answers
	// 6. Create a variable to count our correct answers
	correctAns := 0

	// Channel to receive user answers
	answerCh := make(chan string)

	// Loop through problems
	// 7. Loop through the problems, print the questions and options we'll accept the answers
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, p.q)
		fmt.Printf("Options: %s\n", strings.Join(p.options, ", "))


		// Launch goroutine to handle user input
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer // Send user answer to channel
		}()

		select {
		case <-tObj.C: // Timer expired
			fmt.Println("Time's up!")
			break problemLoop
		case userAnswer := <-answerCh:
			userAnswer = strings.TrimSpace(userAnswer)
			if userAnswer == p.a {
				fmt.Println("Correct!")
				correctAns++
			}else {
				fmt.Printf("Incorrect. The correct answer is: %s\n", p.a)
			}

			// Close channel if it's the last problem
			if i == len(problems)-1 {
				close(answerCh)
			}
		}
	}

	// Output results
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit")
	fmt.Scanln() // Wait for user to press Enter
}

// exit prints an error message and exits the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

