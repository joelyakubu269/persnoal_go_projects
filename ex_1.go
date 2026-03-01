package main

import (
	"fmt"
)

func main() {
	students := []map[string]interface{}{}
	students = append(students, map[string]interface{}{"name": "joyce", "score": 80, "grade": "A"})
	students = append(students, map[string]interface{}{"name": "joel", "score": 90, "grade": "A"})
	students = append(students, map[string]interface{}{"name": "Ahubi", "score": 69, "grade": "B"})

	for _, v := range students {
		name := v["name"].(string)
		grade := v["grade"].(string)
		score := v["score"].(int)
		fmt.Printf("%s scored %d and got grade % s", name, score, grade)
	}
}
