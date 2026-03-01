package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("RESULTS COMPLIER")
	reader := bufio.NewReader(os.Stdin)
	type student struct {
		Name  string
		Grade string
		Score int
	}
	var students []student
	for {
		var student student
		fmt.Print("What is Your Name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Name cannot be empty")
			continue
		}
		hasdigits := false
		for _, r := range input {
			if r >= '0' && r <= '9' {
				hasdigits = true
				break
			}
		}
		if hasdigits {
			fmt.Println("Name cannot have digits")
			continue
		}
		student.Name = input

		fmt.Println("What did you score: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Score cannot be empty")
			continue
		}
		hasLetters := false
		for _, r := range input {
			if r < '0' || r > '9' {
				hasLetters = true
				break
			}
		}
		if hasLetters {
			fmt.Println("score can only be a number")
			continue
		}
		scored, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error...in converting score")
			continue
		}
		if scored < 0 {
			fmt.Println("you are not that dumb. or are you!?")
			continue
		}
		student.Score = scored
		switch {
		case scored >= 120:
			fmt.Println("Excellent results,keep it up")
			student.Grade = "A"
		case scored >= 100:
			fmt.Println("Good results but can do better")
			student.Grade = "B"
		case scored >= 80:
			fmt.Println("Fair results")
			student.Grade = "C"
		default:
			fmt.Println("you will learn the hard way")
			student.Grade = "F"
		}
		students = append(students, student)
		fmt.Println("Do you wish to add another student? (yes/no): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "yes" {
			break
		}
	}
	for _, s := range students {
		fmt.Printf("%s scored %d and got %s\n", s.Name, s.Score, s.Grade)

	}
}
