package main

import (
	"fmt"
)

func sumHourGlass(ar [][]int) int {
	n := len(ar)
	i := 0
	sum := 0
	for {
		j := 0
		if i+2 >= n {
			break
		}
		for {
			if j+2 >= n {
				break
			}
			hourGlass := ar[i][j] + ar[i][j+1] + ar[i][j+2] + ar[i+1][j] + ar[i+1][j+1] + ar[i+1][j+2] + ar[i+1][j+1]
			sum += hourGlass
			j++
		}
		i++
	}
	return sum
}

func main() {
	var question = [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}
	answer := sumHourGlass(question)
	fmt.Println(answer)
}
