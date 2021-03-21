package main

import (
	"fmt"
	"strings"
)

func longestWords(w string) string {
	var wordArray []string = strings.Split(w, " ")
	var lenTemp int = 0
	var indexTemp int
	for x := len(wordArray) - 1; x >= 0; x-- {
		if len(wordArray[x]) > lenTemp {
			lenTemp = len(wordArray[x])
			indexTemp = x
		}
	}
	return wordArray[indexTemp]
}

func main() {
	question := "NEVER ODD OR EVEN PIZZAHUT"
	answer := longestWords(question)
	fmt.Println(answer)
}
