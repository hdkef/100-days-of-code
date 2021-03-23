package main

import (
	"fmt"
	"strings"
)

func confirmEnding(a string, b string) bool {
	trimmed := strings.ReplaceAll(a, " ", "")
	return string(trimmed[len(trimmed)-1]) == b
}

func main() {
	question := "In The Middle"
	answer := confirmEnding(question, "e")
	fmt.Println(answer)
}
