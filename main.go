package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	initString := strings.ReplaceAll(s, " ", "")
	var reversedString []string
	for x := len(initString) - 1; x >= 0; x-- {
		reversedString = append(reversedString, string(initString[x]))
	}
	return initString == strings.Join(reversedString, "")
}

func main() {
	question := "NEVER ODD OR EVEN"
	answer := checkPalindrome(question)
	fmt.Println(answer)
}
