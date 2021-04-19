package main

import (
	"fmt"
	"reflect"
)

func backOrigin(s1, s2 string) bool {
	var s1Ar []string
	var s2Ar []string
	for i := 0; i < len(s1); i++ {
		if string(s1[i]) == "#" {
			s1Ar = s1Ar[:len(s1Ar)-1]
		} else {
			s1Ar = append(s1Ar, string(s1[i]))
		}
	}
	for i := 0; i < len(s2); i++ {
		if string(s2[i]) == "#" {
			s2Ar = s2Ar[:len(s2Ar)-1]
		} else {
			s2Ar = append(s2Ar, string(s2[i]))
		}
	}

	return reflect.DeepEqual(s1Ar, s2Ar)
}

func main() {
	question := "abcde#####"
	answer := backOrigin(question, "xyzvb####s#")
	fmt.Println(answer)
}
