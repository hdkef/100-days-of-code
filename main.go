package main

import (
	"fmt"
)

func findBig(a []int) int {
	biggestInt := 0
	for x := 0; x <= len(a)-1; x++ {
		if a[x] > biggestInt {
			biggestInt = a[x]
		}
	}
	return biggestInt
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 999999, 1023}
	answer := findBig(question)
	fmt.Println(answer)
}
