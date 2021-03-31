package main

import (
	"fmt"
	"sort"
)

type multiDim [][]int

func (a multiDim) Len() int {
	return len(a)
}

func (a multiDim) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a multiDim) Less(i, j int) bool {
	return a[j][0] < a[i][0]
}

func highestLuck(ar [][]int, k int) int {
	var sortedArray multiDim
	sortedArray = multiDim(ar)
	maxPoints := 0
	m := 0
	sort.Sort(multiDim(sortedArray))
	fmt.Println(sortedArray)
	for x := 0; x <= len(sortedArray)-1; x++ {
		if sortedArray[x][1] != 0 && m < k {
			m++
			maxPoints += sortedArray[x][0]
		} else if sortedArray[x][1] == 0 {
			maxPoints += sortedArray[x][0]
		}
	}
	return maxPoints
}

func main() {
	var question = [][]int{{1, 1}, {2, 1}, {3, 0}, {4, 1}, {5, 0}, {6, 1}}
	answer := highestLuck(question, 2)
	fmt.Println(answer)
}
