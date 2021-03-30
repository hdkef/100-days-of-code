package main

import (
	"fmt"
	"strings"
)

func balanceBracket(ar string) bool {
	a := strings.Split(ar, "")
	var bracket []string
	for x := 0; x <= len(a)-1; x++ {
		if a[x] == "{" || a[x] == "[" || a[x] == "(" {
			bracket = append(bracket, a[x])
		} else {
			if len(bracket) == 0 {
				bracket = []string{"NO"}
				break
			}
			popval := bracket[len(bracket)-1]
			switch a[x] {
			case "[":
				if popval != "]" {
					bracket = append(bracket, "NO")
				}
				break
			case "{":
				if popval != "}" {
					bracket = append(bracket, "NO")
				}
				break
			case "(":
				if popval != ")" {
					bracket = append(bracket, "NO")
				}
				break
			}
			bracket = bracket[:len(bracket)-1]
		}
	}
	return len(bracket) == 0
}

func main() {
	var question = "[{(){}}]()]"
	answer := balanceBracket(question)
	fmt.Println(answer)
}
