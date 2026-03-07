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

	finalRes := []int{}
	for len(res) > 0 {
		minIndex := 0
		for i := 1; i < len(res); i++ {
			if res[i] < res[minIndex] {
				minIndex = i
			}
		}
		finalRes = append(finalRes, res[minIndex])
		res = append(res[:minIndex], res[minIndex+1:]...)
	}
	return finalRes
}
