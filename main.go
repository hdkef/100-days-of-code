package main

import (
	"fmt"
)

func pascalTriangle(k int) [][]int {
	triangle := [][]int{}
	triangle = append(triangle, []int{1})
	for i := 0; i < k; i++ {
		if i <= 0 {
			continue
		} else {
			prevRow := triangle[i-1]
			newRow := []int{1}
			for j := 0; j < i-1; j++ {
				newRow = append(newRow, prevRow[j+1]+prevRow[j])
			}
			newRow = append(newRow, 1)
			triangle = append(triangle, newRow)
		}
	}
	return triangle
}

func main() {
	question := 5
	answer := pascalTriangle(question)
	fmt.Println(answer)
}
