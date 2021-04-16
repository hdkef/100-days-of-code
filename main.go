package main

import (
	"fmt"
	"math"
)

func threeSum(ar []int, k int) []int {
	res := math.MaxInt64
	resAr := []int{}
	for i := 0; i < len(ar); i++ {
		a_pointer := 0
		b_pointer := len(ar) - 1
		for {
			tmp := ar[i] + ar[a_pointer] + ar[b_pointer]
			if a_pointer >= b_pointer {
				break
			}
			if ar[a_pointer]+ar[b_pointer] > k {
				b_pointer--
			} else {
				a_pointer++
			}
			if math.Abs(float64(tmp)-float64(k)) < math.Abs(float64(res)-float64(k)) {
				res = tmp
				resAr = []int{ar[i], ar[a_pointer], ar[b_pointer]}
			}
		}
	}
	return resAr
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	answer := threeSum(question, 19)
	fmt.Println(answer)
}
