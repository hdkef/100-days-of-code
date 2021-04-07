package main

import (
	"fmt"
	"strings"
)

func twoStrings(a1, a2 string) bool {
	for i, _ := range a1 {
		if strings.Contains(a2, string(a1[i])) == true {
			return true
		}
	}
	return false
}

func main() {
	var question string = "hello"
	answer := twoStrings(question, "zzz")
	fmt.Println(answer)
}
