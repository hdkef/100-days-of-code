package main

import (
	"fmt"
	"strings"
)

func titleCase(w string) string {
	wordArray := strings.Split(w, " ")
	var titleCased []string
	for x := 0; x <= len(wordArray)-1; x++ {
		charArray := strings.Split(wordArray[x], "")
		charArray[0] = strings.ToUpper(charArray[0])
		titleCased = append(titleCased, strings.Join(charArray, ""))
	}
	return strings.Join(titleCased, " ")
}

func main() {
	question := "how can i move on when i'm still in love"
	answer := titleCase(question)
	fmt.Println(answer)
}
