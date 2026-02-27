package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var fullName string
	var rage int
	for {
		fmt.Println("USER INFORMATION PROGRAM\n")
		fmt.Println("PLEASE ENTER YOUR FULLNAME: ")
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
			fmt.Println("only alphabets are allowed")
			continue
		}
		if !strings.Contains(input, " ") {
			fmt.Println("enter your last name too")
			continue
		}
		fullName = input
		break
	}
	fmt.Println("hello", fullName)
	for {
		fmt.Println("what is your age: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("age cant be empty or zero")
			continue
		}
		input = strings.ReplaceAll(input, " ", "")
		age, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Age must be a number")
			continue
		}
		if age < 0 {
			fmt.Println("age cannot be less than zero")
			continue
		}
		rage = age
		break
	}
	fmt.Println("hello", fullName, "you age is ", rage)
}
