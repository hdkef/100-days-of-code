package main

import (
	"fmt"
	"math"
)

func lowestAbsDif(ar []int) float64 {
	var min float64 = math.MaxFloat64
	for i := 0; i <= len(ar)-1; i++ {
		for j := 0; j <= len(ar)-1; j++ {
			if j == i {
				continue
			} else {
				dif := math.Abs(float64(ar[j]) - float64(ar[i]))
				if dif < min {
					min = dif
				}
			}
		}
	}
	return min
}

func main() {
	var question = []int{10, 1, 200}
	answer := lowestAbsDif(question)
	fmt.Println(answer)
}
