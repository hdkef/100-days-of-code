package main

import (
	"fmt"
)

func findPath(a []int) []int {
	m := len(a)
	i := -1
	var path []int
	for {
		i += 2
		if i > m {
			break
		}
		fmt.Println("i", i)
		fmt.Println("path", path)
		switch a[i] {
		case 0:
			path = append(path, i)
		case 1:
			i -= 1
			if a[i] == 0 {
				path = append(path, i)
			}
		}
	}
	return path
}

func main() {
	var question = []int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0}
	answer := findPath(question)
	fmt.Println(answer)
}
