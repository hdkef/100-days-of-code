package main

import (
	"fmt"
	"strings"
)

func countValley(ar string) int {
	a := strings.Split(ar, "")
	numval := 0
	alt := 0
	for x := 0; x <= len(a)-1; x++ {
		if a[x] == "U" {
			if alt == -1 {
				numval++
			}
			alt++
		} else if a[x] == "D" {
			alt--
		} else {
			fmt.Println("err")
		}
	}
	return numval
}

func main() {
	var question = "DDUUDDUUDDUU"
	answer := countValley(question)
	fmt.Println(answer)
}
