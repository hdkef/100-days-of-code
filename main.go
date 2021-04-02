package main

import (
	"fmt"
	"strings"
)

func alternateChar(s string) string {
	arChar := strings.Split(s, "")
	var charStack []string
	for x, _ := range arChar {
		if x == len(arChar)-1 {
			charStack = append(charStack, arChar[x])
			break
		} else if arChar[x] == arChar[x+1] && x < len(arChar) {
			continue
		} else {
			charStack = append(charStack, arChar[x])
		}
	}

	return strings.Join(charStack, "")
}

func main() {
	var question string = "AAABBBCCCDDDEEE"
	answer := alternateChar(question)
	fmt.Println(answer)
}
