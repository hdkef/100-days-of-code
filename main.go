package main

import (
	"fmt"
	"strings"
)

func palindromeNumber(a int) bool {
	stringify := fmt.Sprint(a)
	charArr := strings.Split(stringify, "")
	var reverseArr []string
	for x := len(charArr) - 1; x >= 0; x-- {
		reverseArr = append(reverseArr, charArr[x])
	}
	return strings.Join(reverseArr, "") == stringify
}

func main() {
	question := 121
	answer := palindromeNumber(question)
	fmt.Println(answer)
}
