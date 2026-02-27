package main

import (
	"fmt"
	"os"
	"strings"

	"bufio"
	"strconv"
)

func main() {
	var fullName string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("USER INFORMATION PROGRAM\n \n")

		fmt.Println("ENTFER YOUR NAME: ")
		input, _ := reader.ReadString(('\n'))
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Name cannot be empty")
			continue
		}

		hasnum := false
		for _, r := range input {
			if r >= '0' && r <= '9' {
				hasnum = true
				break
			}
		}
		if hasnum {
			fmt.Println("Write in alphabet format only")
			continue
		}
		if !strings.Contains(input, " ") {
			fmt.Println("Enter.. your last name too")
			continue

		}
		fullName = input
		break
	}
	fmt.Printf("Hello % s, you are what is your age: ", fullName)
	var age int
	var err error
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Age cannot be empty")
			continue
		}
		input = strings.ReplaceAll(input, " ", "")
		age, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error Age cant be in alphabetical format")
			continue
		}
		if age < 0 {
			fmt.Println("age cannot be negative")
			continue
		}
		break
	}
	fmt.Println(fullName, "your age is ", age)

}
