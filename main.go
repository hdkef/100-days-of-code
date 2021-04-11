package main

import (
	"fmt"
)

func twoSum(ar []int, target int) []int {
	a_pointer := 0
	b_pointer := len(ar) - 1
	sum := ar[a_pointer] + ar[b_pointer]
	for {
		sum = ar[a_pointer] + ar[b_pointer]
		if sum < target {
			a_pointer++
		} else if sum > target {
			b_pointer--
		} else {
			return []int{ar[a_pointer], ar[b_pointer]}
		}
	}
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8}
	answer := twoSum(question, 8)
	fmt.Println(answer)
}
