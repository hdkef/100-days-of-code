package main

import (
	"fmt"
	"math"
)

func diagonalDiff(m [][]int) int {
	leftToRight := 0
	rightToLeft := 0
	i := 0
	for {
		if i > len(m)-1 {
			break
		}
		j := len(m) - 1 - i
		leftToRight += m[i][i]
		rightToLeft += m[i][j]
		fmt.Println(leftToRight, rightToLeft)
		i++
	}
	return int(math.Abs(float64(leftToRight) - float64(rightToLeft)))
}

func main() {
	var question = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	answer := diagonalDiff(question)
	fmt.Println(answer)
}
