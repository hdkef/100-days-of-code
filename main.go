package main

import (
	"fmt"
	"strings"
)

func reverseOnlyLetter(s string) string {
	var letterOnly []string
	var reversed []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "-" {
			letterOnly = append(letterOnly, string(s[i]))
		}
	}
	for j := 0; j < len(s); j++ {
		if string(s[j]) != "-" {
			reversed = append(reversed, letterOnly[len(letterOnly)-1:]...)
			letterOnly = letterOnly[:len(letterOnly)-1]
		} else {
			reversed = append(reversed, string(s[j]))
		}
	}
	return strings.Join(reversed, "")
}

func main() {
	question := "abc-defg-hijk"
	answer := reverseOnlyLetter(question)
	fmt.Println(answer)
}
