package main

import (
	"fmt"
)

func leftRotation(ar []int, k int) []int {
	if k-1 > len(ar) {
		return ar
	}
	var newAr []int
	newAr = ar[k:]
	newAr = append(newAr, ar[:k]...)
	return newAr
}

func main() {
	var question = []int{1, 2, 3, 4, 5}
	answer := leftRotation(question, 2)
	fmt.Println(answer)
}
