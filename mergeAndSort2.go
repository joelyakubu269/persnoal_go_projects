package main

import "fmt"

func main() {
	var a = []int{2, 3, 4, 5}
	var b = []int{6, 74, 34, 3, 2}
	fmt.Println(mergeSlice(a, b))
}
func mergeSlice(s, t []int) []int {
	res := []int{}
	res = append(res, s...)
	res = append(res, t...)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[j] < res[i] {
				res[j], res[i] = res[i], res[j]
			}
		}
	}
	return res
}
