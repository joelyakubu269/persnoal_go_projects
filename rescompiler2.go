package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//	reader := bufio.NewReader(os.Standin)
	res := []map[string]interface{}{}
	//res = append(res, student)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("RESULTS COMPILER")
	for {
		student := map[string]interface{}{}
		for {

			fmt.Println("What is your name: ")
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
				fmt.Println("name can contain only letters")
				continue
			}
			student["name"] = input

			break
		}
		for {
			fmt.Println("What did you score: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("your brain maybe empty but not ur score at least")
				continue
			}
			hasletters := false
			for _, r := range input {
				if r < '0' || r > '9' {
					hasletters = true
					break
				}
			}
			if hasletters {
				fmt.Println("Score cannot contain letters")
				continue

			}
			scored, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error...cannot convert to int")
				continue
			}
			if scored < 0 {
				fmt.Println("score cannot be negative")
				continue
			}
			student["score"] = scored
			switch {
			case scored >= 120:
				fmt.Println("excellento,bravos signor")
				student["grade"] = "A"
			case scored >= 100:
				fmt.Println("keep scaling up, there is work to be done")
				student["grade"] = "C"
			default:
				fmt.Println("get your self right ")
				student["grade"] = "F"
			}
			break
		}

		res = append(res, student)
		fmt.Println("Do you want to add another?(yes/no): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "yes" {
			break
		}

	}
	for _, r := range res {
		fmt.Printf("%s scored %d and got an %s\n", r["name"].(string), r["score"].(int), r["grade"].(string))
	}
}
// in this method i used nested for loops instead of one like the former
// advantages of using this nested for loop
// * Two loop version only repeats the part that failed-* cleaner ux for user
// * appends after all fields are validated.always safe no empty maps in the slice
// no count needed
