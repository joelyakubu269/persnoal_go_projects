package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("THIS IS A SIMPLE RESULT COMPILER")
	var students = []map[string]interface{}{}

	reader := bufio.NewReader(os.Stdin)
	for {
		var student = map[string]interface{}{}
		for {

			fmt.Println("WHAT IS YOUR NAME: ")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				fmt.Println("ERROR in reading file")
				return
			}
			if input == "" {
				fmt.Println("Name cannot be empty")
				continue
			}
			notAlpha := false
			for _, r := range input {
				if !unicode.IsLetter(r) && r != ' ' && r != '-' {
					notAlpha = true
					break
				}
			}
			if notAlpha {
				fmt.Println("Enter only alphabets")
				continue
			}
			if !strings.Contains(input, " ") {
				fmt.Println("Enter your full name")
				continue
			}
			words := strings.Fields(input)
			for i, w := range words {
				words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
			}
			input = strings.Join(words, " ")
			student["name"] = input

			break
		}

		for {
			fmt.Println("WHAT DID YOU SCORE: ")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				fmt.Println("There was an error this file")
				return
			}
			if input == "" {
				fmt.Println("Score cannot be empty")
				continue
			}

			Score, err := strconv.Atoi(input)
			if err != nil || Score < 0 || Score > 100 {
				fmt.Println("Enter a valid score (0-100)")
				continue
			}
			student["score"] = Score

			switch {
			case Score >= 90:
				fmt.Println("Great results keep it up")
				student["Grade"] = "A"
			case Score >= 80:
				fmt.Println("Good results")
				student["Grade"] = "B"
			case Score >= 70:
				fmt.Println("Mediocrity gets you no where")
				student["Grade"] = "C"
			default:
				fmt.Println("Pick yourself up lad")
				student["Grade"] = "F"

			}

			break

		}
		students = append(students, student)
		fmt.Println("Do you want to add another student (yes/no): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		if input != "yes" {
			break
		}

	}

	fmt.Println("\nFINAL RESULTS:")
	for i, s := range students {
		fmt.Printf("%d. %s scored %d → Grade: %s\n", i+1, s["name"].(string), s["score"].(int), s["Grade"].(string))
	}
}
