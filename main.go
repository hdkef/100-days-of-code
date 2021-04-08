package main

import (
	"fmt"
)

func numIsland(arr [][]int) int {
	islandNum := 0
	for i, _ := range arr {
		for j, _ := range arr[i] {
			if arr[i][j] == 1 {
				islandNum++
				findAndDestroy(&arr, i, j)
			}
		}
	}
	return islandNum
}

func findAndDestroy(arr *[][]int, i int, j int) {
	if i < 0 || i > len(*arr)-1 || j < 0 || j > len((*arr)[i])-1 || (*arr)[i][j] == 0 {
		return
	}
	(*arr)[i][j] = 0
	findAndDestroy(arr, i+1, j)
	findAndDestroy(arr, i-1, j)
	findAndDestroy(arr, i, j+1)
	findAndDestroy(arr, i, j-1)
}

func main() {
	question := [][]int{{1, 1, 1, 1}, {1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 1, 1}, {1, 0, 0, 0}}
	answer := numIsland(question)
	fmt.Println(answer)
}
