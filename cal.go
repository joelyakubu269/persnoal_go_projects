package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val1 string
	var val2 string
	var operator string

	for {

		fmt.Println("A SIMPLE CALCULATOR THAT ONLY MULTIPLIES,DIVIDES AND ADDS")
		fmt.Println("Enter Your Values : ")
		fmt.Scan(&val1, &operator, &val2)
		if val1 == "" || val2 == "" {
			fmt.Println("Error...values cannot be empty")

		}
		if operator != "*" && operator != "+" && operator != "/" && operator != "-" {
			fmt.Println("operation not allowed")
			continue

		}
		val1N, err := strconv.Atoi(val1)
		if err != nil {
			fmt.Println("value must be a number")
			continue
		}
		val2N, err := strconv.Atoi(val2)
		if err != nil {
			fmt.Println("value must be a number")
			continue
		}

		if operator == "" {
			fmt.Println("place an operator either + or * or divides")
			continue

		}
		if operator == "+" {
			fmt.Println(val1N + val2N)
		}
		if operator == "/" {
			if val2N == 0 {
				fmt.Println("You are not ramanujan you cant know infinity")
				continue
			}
			{
				fmt.Println(val1N / val2N)
			}
		}

		if operator == "*" {
			fmt.Println(val1N * val2N)
		}
		if operator == "-" {
			fmt.Println(val1N - val2N)
		}
		break
	}

}
