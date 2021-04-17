package main

import (
	"fmt"
	"sort"
)

type cards []int

func (c cards) Len() int {
	return len(c)
}

func (c cards) Swap(a, b int) {
	c[a], c[b] = c[b], c[a]
}

func (c cards) Less(a, b int) bool {
	return c[a] < c[b]
}

func handOfStraight(ar []int, k int) [][]int {
	cards := cards(ar)
	sort.Sort(cards)
	groupModulus := len(ar) % k
	if groupModulus != 0 {
		return [][]int{}
	} else {
		var groupResult [][]int
		totalGroup := len(ar) / k
		for i := 0; i < len(ar); i += totalGroup {
			var grouped []int
			for j := i; j <= i+totalGroup-1; j++ {
				grouped = append(grouped, ar[j])
			}
			groupResult = append(groupResult, grouped)
		}
		return groupResult
	}
}

func main() {
	question := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	answer := handOfStraight(question, 4)
	fmt.Println(answer)
}
