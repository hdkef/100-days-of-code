package main

import (
	"fmt"
)

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var firstArr []byte = []byte(a)
	var secondArr []byte = []byte(b)
	var firstSum byte
	var secondSum byte
	for x := 0; x < len(firstArr); x++ {
		firstSum += firstArr[x]
		secondSum += secondArr[x]
	}
	return firstSum == secondSum
}

func main() {
	var question string = "PEOPLE"
	answer := isAnagram(question, "PLEOEP")
	fmt.Println(answer)
}
