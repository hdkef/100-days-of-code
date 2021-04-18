package main

import (
	"fmt"
	"strings"
)

func backOrigin(path string) bool {
	pathAr := strings.Split(path, "")
	var U []int
	var D []int
	var R []int
	var L []int
	for i := 0; i < len(pathAr); i++ {
		if pathAr[i] == "U" {
			U = append(U, 1)
		} else if pathAr[i] == "D" {
			D = append(D, 1)
		} else if pathAr[i] == "R" {
			R = append(R, 1)
		} else if pathAr[i] == "L" {
			L = append(L, 1)
		} else {
			return false
		}
	}

	return len(U) == len(D) && len(R) == len(L)
}

func apiCall(s int) bool {
	if s >= 17 {
		return true
	} else {
		return false
	}
}

func main() {
	question := "UUUULLLLDDDDRRRRD"
	answer := backOrigin(question)
	fmt.Println(answer)
}
