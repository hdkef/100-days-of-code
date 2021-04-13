package main

import (
	"fmt"
	"math"
)

func binarySearch(ar []int, k int) int {
	first_i := 0
	last_i := len(ar) - 1
	middle_i := math.Ceil((float64(first_i) + float64(last_i)) / 2)

	for {
		if ar[int(middle_i)] < k {
			if ar[last_i] == k {
				return last_i
			}
			first_i = int(middle_i)
			middle_i = math.Ceil((float64(first_i) + float64(last_i)) / 2)
		} else if ar[int(middle_i)] > k {
			if ar[first_i] == k {
				return first_i
			}
			last_i = int(middle_i)
			middle_i = math.Ceil((float64(first_i) + float64(last_i)) / 2)
		} else if ar[int(middle_i)] == k {
			return int(middle_i)
		}
	}
}

func main() {
	question := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	answer := binarySearch(question, 7)
	fmt.Println(answer)
}
