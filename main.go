package main

import (
	"fmt"
)

func findDuplicate(ar []int) bool {
	tmpMap := make(map[int]int)
	for x := 0; x < len(ar); x++ {
		_, y := tmpMap[ar[x]]
		if y == false {
			tmpMap[ar[x]] = ar[x]
		} else {
			return true
		}
	}
	return false
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 9}
	answer := findDuplicate(question)
	fmt.Println(answer)
}
