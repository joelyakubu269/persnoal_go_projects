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
	fmt.Println("SEMESTER RESULTS COMPLIER")
	res := []map[string]interface{}{}

	//var name string
	var score int
	var err error
	var count = 0
	for {
		res = append(res, map[string]interface{}{})
		fmt.Println("WHAT IS YOUR NAME: ")
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
			fmt.Println("Your name cannot contain digits")
			continue
		}
		res[count]["name"] = input

		fmt.Println("What did you score in maths: ")
		input, _ = reader.ReadString('\n') // input was not redecleared so it can overwrite the former value
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("score cannot be empty")
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
			fmt.Println("score cannot contain letters")
			continue
		}
		input = strings.ReplaceAll(input, " ", "")
		score, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error in coverting to score")
			continue
		}
		if score < 0 {
			fmt.Println("score cannot be negative")
		}
		switch {
		case score >= 120:
			res[count]["grade"] = "A"
			fmt.Printf("Great result %s,keep it up\n", res[count]["name"])
		case score >= 100:
			res[count]["grade"] = "B"
			fmt.Printf("Good result %s,you can do better\n", res[count]["name"])
		case score >= 80:
			res[count]["grade"] = "C"
			fmt.Printf("This is a fair result %s, buckle up", res[count]["name"])
		default:
			res[count]["grade"] = "F"
			fmt.Println("Kanka ba daya ba")
		}
		count++
		fmt.Println("do you want to add another student? (y/n): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" {
			break
		}

	}
	fmt.Println(res)
	fmt.Printf("%s scored")
}

/* for {
    // Create a new map for this student
    student := map[string]interface{}{}

    // Fill the student info
    fmt.Print("WHAT IS YOUR NAME: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    student["name"] = input

    // ... handle score and grade ...

    // Append this student map to the results slice
    res = append(res, student)

    // Ask to continue or break
} Notice now we dont need count because append automatically adds the new map to the slice
*/
