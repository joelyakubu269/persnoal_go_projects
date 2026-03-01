package main

import (
	"fmt"
)

func main() {
	fmt.Print("SEMESTER SCORE CARD")
	res := []map[string]interface{}{}                                                     // declaring of the slice of maps
	res = append(res, map[string]interface{}{"name": "joyce", "score": 100, "grade": ""}) // map is initialized and values can be appended
	res = append(res, map[string]interface{}{"name": "joel", "score": 120, "grade": ""})
	res = append(res, map[string]interface{}{"name": "Ahubi", "score": 80, "grade": ""})
	fmt.Println(res)
	for _, v := range res {
		name := v["name"].(string)
		scoreValue, ok := v["score"]

		if !ok {
			fmt.Printf("This %s does not have a score\n", name)
		}
		score := scoreValue.(int)
		switch {
		case score >= 120:
			v["grade"] = "A"
		case score >= 100:
			v["grade"] = "B"
		case score >= 80:
			v["grade"] = "C"
		}
		fmt.Printf("%s scored %d and the grade is %v", name, score, v["grade"])
	}
	fmt.Println(res)
}
