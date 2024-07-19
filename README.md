# Golang-Simple-quiz-application

Welcome to the Golang-Simple-quiz-application, a simple yet efficient quiz tool built with Go. This application reads quiz questions from a CSV file, presents them to the user, and times the entire quiz session. It's designed to be straightforward and easy to use, perfect for quick tests or learning exercises.

## Features

- **CSV Parsing**: Load quiz questions, options, and answers from a CSV file.
- **Customizable Timer**: Set the quiz duration with a command-line flag.
- **Real-Time Answer Checking**: Get immediate feedback on your answers.
- **Concurrent Processing**: Handles user input and timing simultaneously using goroutines.

## Prerequisites
- [Go](https://golang.org/doc/install) installed on your machine.

## Running the application 

 Create a CSV file (e.g., `quiz.csv`) with the following format:
   ```csv
   question,correct_answer,optionA,optionB,optionC,optionD
   45+30,75,A) 70,B) 75,C) 80,D) 85
   88+32,120,A) 105,B) 110,C) 115,D) 120
  ```

Open terminal and type the following command:
  ```bash
   go run main.go -f quiz.csv -t 30
  ```

- **-f** specifies the path to the CSV file.
- **-t** sets the quiz duration in seconds.

## Example

``` bash
$ go run main.go -f quiz.csv -t 30
Problem 1: 45+30
Options: A) 70, B) 75, C) 80, D) 85
Your answer: B
Correct!

Problem 2: 88+32
Options: A) 105, B) 110, C) 115, D) 120
Your answer: D
Correct!

Time's up!
Your result is 2 out of 5
Press enter to exit
```



