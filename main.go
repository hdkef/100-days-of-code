package main

import (
	"fmt"
	"strings"
)

func balanceBracket(ar string) bool {
	a := strings.Split(ar, "")
	var bracket []int
	for x := 0; x <= len(a)-1; x++ {
		if a[x] == "{" || a[x] == "[" || a[x] == "(" {
			bracket = append(bracket, 1)
		} else {
			if len(bracket) == 0 {
				bracket = []int{1}
				break
			}
			bracket = bracket[0 : len(bracket)-1]
		}
	}
	return len(bracket) == 0
}

func main() {
	var question = "[{(){}}]"
	answer := balanceBracket(question)
	fmt.Println(answer)
}
