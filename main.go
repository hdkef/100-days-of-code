package main

import (
	"fmt"
)

func countPair(ar []int) int {
	mapPair := make(map[int]int)
	numPair := 0
	for _, x := range ar {
		_, found := mapPair[x]
		if found == false {
			mapPair[x] = x
		} else {
			delete(mapPair, x)
			numPair++
		}
	}
	return numPair
}

func main() {
	var question = []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7}
	answer := countPair(question)
	fmt.Println(answer)
}
