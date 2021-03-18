# 100 Days Of Code - Log

### Day 0: Mar 18th 2021

**Today's Progress**: 
i learned how to reverse string using javascript 

and i learned how to reverse string using golang



**Thoughts:** :
there are some differences between a for loop in javascript and golang.
for example, for let x := 3; x--; x>0{}
javascript will substract 3 with 1 first, so x = 2 is the first go in the body of a function, but
Golang will NOT substract 3 with 1 first, so x = 3 is the first go in the body of a function.

in order to work the loop, we must consider the length of an arrays / slices and indexes (especially
in golang ). example : len(array) - 1

**Link to work:**:
this is the code :

javascript

function reverseString(s){
  let reversedString = []
  let stringArray = s.split(" ")
  for (let x = stringArray.length -1;x--;x>0){
  	let temp = []
  	for(let y = stringArray[x].length - 1;y--;y>0){
      temp.push(stringArray[x][y])
    }
    reversedString.push(temp.join(""))
    temp = []
  }
  return reversedString.join("	")
}

let question = "HELLO WORLD"
let reversedQuestion = reverseString(question)
console.log(reversedQuestion)

golang

package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	stringArray := strings.Split(s, " ")
	var reversedString []string
	for x := len(stringArray) - 1; x >= 0; x-- {
		var temp []string
		for y := len(stringArray[x]) - 1; y >= 0; y-- {
			temp = append(temp, string(stringArray[x][y]))
		}
		reversedString = append(reversedString, strings.Join(temp, ""))
		temp = nil
	}
	return strings.Join(reversedString, "	")
}

func main() {
	var question string = "Hello playground"
	var answer string = reverseString(question)
	fmt.Println(answer)
}

