# 100 Days Of Code - Log

### Day 0: Mar 18th 2021

**Today's Progress**: 
i learned how to reverse string using javascript 

and i learned how to reverse string using golang



**Thoughts:** :

first we need to split words between spaces, then rearrange words.

**Link to work:**:
this is the code :

javascript

function reverseString(s){
    let reversedString = []
    let stringArray = s.split(" ")
    for (let x = stringArray.length - 1;x>=0;x--){
        let temp = []
        for(let y = stringArray[x].length - 1;y>=0;y--){
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

### Day 1: 19th Mar 2021

**Today's Progress:**
i learned how to factorialize using javascript and golang with two methods.
Recursive function and standard.

**Thoughts:**
recursive -> we return the value * function(value - 1) and if value == 0 return 1
non recursive -> we do for loop with decremental i then value = value * i

**Link to Work:**

javascript recursive method

function factorialize(i){
	if (i == 0){
  	return 1
  }
  return i * factorialize(i-1)
}

let question = 10
let answer = factorialize(5)
console.log(answer)

golang recursive method

package main

import "fmt"

func factorialize(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorialize(i-1)
}

func main() {
	question := 5

	var answer = factorialize(question)

	fmt.Println(answer)
}


javascript non recursive method

function factorialize(i){
let tmp = 1
for (let x = i;x>0;x--){
console.log(x)
  	tmp = tmp * x
}
return tmp
}

let question = 10
let answer = factorialize(question)
console.log(answer)


golang non recursive method

package main

import "fmt"

func factorialize(i int) int {
	var tmp = 1
	for x := i; x > 0; x-- {
		tmp = tmp * x
	}
	return tmp
}

func main() {
	question := 5
	answer := factorialize(question)
	fmt.Println(answer)
}


