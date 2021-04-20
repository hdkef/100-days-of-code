package main

import (
	"fmt"
)

func jewelsAndStones(j, s string) int {
	var count int
	for i := 0; i < len(j); i++ {
		for k := 0; k < len(s); k++ {
			if j[i] == s[k] {
				count++
			}
		}
	}
	return count
}

func main() {
	question := "abc"
	answer := jewelsAndStones(question, "aaabbbccc")
	fmt.Println(answer)
}
