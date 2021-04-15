package main

import (
	"fmt"
	"math"
)

func containerWithMostWater(ar []int) float64 {
	tmp := 0.0
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[i] < ar[j] {
				size := math.Abs(float64(i)-float64(j)) * float64(ar[i])
				if size > tmp {
					tmp = size
				}
			} else {
				size := math.Abs(float64(i)-float64(j)) * float64(ar[j])
				if size > tmp {
					tmp = size
				}
			}
		}
	}
	return tmp
}

func main() {
	question := []int{1, 99999, 2222, 3, 3333, 5, 7, 90}
	answer := containerWithMostWater(question)
	fmt.Println(answer)
}
