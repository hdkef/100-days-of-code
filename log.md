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

### Day 3: 20th Mar 2021

**Today's Progress:**
i learned how to check palindrome using javascript and golang.

**Thoughts:**
basically just replace space with none (eliminate spaces)then reverse the string using for loop,
push, and joined again, then check the initial string and reversed string. 

**Link to Work:**

javascript

function checkPalindrome(s){
    let initString = s.replace(/\s/g, '')
    let reverseString = []

    for (let x = initString.length - 1;x>=0;x--){
        reverseString.push(initString[x])
    }

    return initString == reverseString.join("")
}

let question = "NEVER ODD OR EVEN"
let answer = checkPalindrome(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	initString := strings.ReplaceAll(s, " ", "")
	var reversedString []string
	for x := len(initString) - 1; x >= 0; x-- {
		reversedString = append(reversedString, string(initString[x]))
	}
	return initString == strings.Join(reversedString, "")
}

func main() {
	question := "NEVER ODD OR EVEN"
	answer := checkPalindrome(question)
	fmt.Println(answer)
}

### Day 4: 21st Mar 2021

**Today's Progress:**
i learned how to find longest words using javascript and golang.

**Thoughts:**
basically just split the sentences using Split assign to wordArray, for loop the len and change longest length to lenTemp and change longest index to indexTemp. Finally return wordArray at index of indexTemp

**Link to Work:**

Javascript

function longestWords(w){
    let wordArray = w.split(" ")
    let lenTemp = 0
    let indexTemp = 0
    for (let x = wordArray.length-1;x>=0;x--){
        if (wordArray[x].length > lenTemp){
            lenTemp = wordArray[x].length
            indexTemp = x
        }
    }
    return wordArray[indexTemp]
}

let question = "NEVER ODD OR EVEN PTERODUCTYL"
let answer = longestWords(question)
console.log(answer)


Golang

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


### Day 5: 22nd Mar 2021

**Today's Progress:**
i learned how to titlecase sentences using javascript and golang.

**Thoughts:**
basically just split the sentences using Split assign to wordArray, for loop the len and create
charArray then charArray[0].toUpperCase(), after that push to TitleCased (is an array). Finally
return TitleCased.join(" )

**Link to Work:**

javascript

function titleCase(w){
    let wordArray = w.split(" ")
    let titleCased = []
    for (let x = 0;x<=wordArray.length-1;x++){
        charArray = wordArray[x].split("")
        charArray[0] = charArray[0].toUpperCase()
        titleCased.push(charArray.join(""))
    }
    return titleCased.join(" ")
}

let question = "how can i move on when i'm still in love"
let answer = titleCase(question)
console.log(answer)

golang

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

### Day 6: 23rd Mar 2021

**Today's Progress:**
i learned how to find biggest integer using javascript and golang.

**Thoughts:**
basically just for loop and check if a[i] > biggestInt and return biggestInt after looping

**Link to Work:**

javascript

function findBig(a){
    let biggestInt = 0
    try{
        for (let x = 0;x<=a.length-1;x++){
            if (a[x] > biggestInt){
                biggestInt = a[x]
            }
        }
        return biggestInt
    }catch(err){
        console.log(err)
    }
}

let question = [2,1,2,34,787,99999]
let answer = findBig(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func findBig(a []int) int {
	biggestInt := 0
	for x := 0; x <= len(a)-1; x++ {
		if a[x] > biggestInt {
			biggestInt = a[x]
		}
	}
	return biggestInt
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 999999, 1023}
	answer := findBig(question)
	fmt.Println(answer)
}


### Day 7: 24th Mar 2021

**Today's Progress:**
i learned how to confirm ending of string using javascript and golang.

**Thoughts:**
basically just access last index (len - 1) of trimmed (replace " ", "") string equal second argument

**Link to Work:**

javascript

function confirmEnding(a,b){
    return a.replace(" ","").substr(-1,1).toUpperCase() == b.toUpperCase()
}

let question = "NOTHING BUT NOTHING"
let answer = confirmEnding(question,"g")
console.log(answer)

golang

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

### Day 8: 25th Mar 2021

**Today's Progress:**
i learned how to find diff of diagonal of matrix (multi dim array) using javascript and golang.

**Thoughts:**
items of left to right diagonal is basically access array with the same index [i][i]
and right to left is [i][j] and j is decrement of the index length. At the end return the math.Abs(ltr - rtl) to make the dif positive

**Link to Work:**

javascript

function diagonalDiff(m){
    var leftToRight = 0
    var rightToLeft = 0
    let i = 0
    while (i <= m.length - 1){
        leftToRight += m[i][i]
        console.log(m[i][i])
        let j = m.length - 1 - i
        rightToLeft += m[i][j]
        console.log(m[i][j])
        i++
        console.log(rightToLeft,leftToRight)
    }
    return Math.abs(leftToRight - rightToLeft)
}

let question = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]
let answer = diagonalDiff(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func diagonalDiff(m [][]int) int {
	leftToRight := 0
	rightToLeft := 0
	i := 0
	for {
		if i > len(m)-1 {
			break
		}
		j := len(m) - 1 - i
		leftToRight += m[i][i]
		rightToLeft += m[i][j]
		fmt.Println(leftToRight, rightToLeft)
		i++
	}
	return int(math.Abs(float64(leftToRight) - float64(rightToLeft)))
}

func main() {
	var question = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	answer := diagonalDiff(question)
	fmt.Println(answer)
}

### Day 9: 26th Mar 2021

**Today's Progress:**
i learned how to solve hackerrank jump on the clouds using javascript and golang.

**Thoughts:**
i need to be consider the placement of break statement and when should we use > or >=

**Link to Work:**

javascript

function findPath(a){
    let pLen = a.length
    let i = -1
    let path = []
    while (i < pLen){
        i+=2
        switch (a[i]){
            case 0:
                path.push(i)
                break
            case 1:
                i-=1
                if (a[i] == 0){
                    path.push(i)
                }
                break
        }
    }
    return path
}

let question = [0,1,0,0,1,0,0,1,0,0,1,0]
let answer = findPath(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func findPath(a []int) []int {
	m := len(a)
	i := -1
	var path []int
	for {
		i += 2
		if i > m {
			break
		}
		fmt.Println("i", i)
		fmt.Println("path", path)
		switch a[i] {
		case 0:
			path = append(path, i)
		case 1:
			i -= 1
			if a[i] == 0 {
				path = append(path, i)
			}
		}
	}
	return path
}

func main() {
	var question = []int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0}
	answer := findPath(question)
	fmt.Println(answer)
}

### Day 11: 28th Mar 2021

**Today's Progress:**
i learned how to examine balanced bracketusing javascript and golang.

**Thoughts:**
basically push if left bracket and pop if right bracket, don't forget to check if we can't pop a zero len array.
Golang use slice to pop btw
Check popped value, if its not the corresponding left bracket return NO

**Link to Work:**

javascript

function balanceBracket(ar){
    let a = ar.split("")
    let bracketStack = []
    for (let x = 0;x <= a.length - 1;x++){
        if (a[x] == '[' || a[x] == '{' || a[x] == '('){
            bracketStack.push(a[x])
        }
        else{
            if (bracketStack.length == 0){
                bracketStack = ["NO"]
                break
            }
            else{
                let popval = bracketStack.pop()
                console.log(popval)
                switch (a[x]){
                    case "}":
                        if (popval != "{"){
                            bracketStack = ["NO"]
                        }
                        break
                    case "]":
                        if (popval != "["){
                            bracketStack = ["NO"]
                        }
                        break
                    case ")":
                        if (popval != "("){
                            bracketStack = ["NO"]
                        }
                        break
                }
            }
        }
    }
    console.log(bracketStack)
    return bracketStack.length == 0
}

let question = "[{(){}}][]"
let answer = balanceBracket(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"strings"
)

func balanceBracket(ar string) bool {
	a := strings.Split(ar, "")
	var bracket []string
	for x := 0; x <= len(a)-1; x++ {
		if a[x] == "{" || a[x] == "[" || a[x] == "(" {
			bracket = append(bracket, a[x])
		} else {
			if len(bracket) == 0 {
				bracket = []string{"NO"}
				break
			}
			popval := bracket[len(bracket)-1]
			switch a[x] {
			case "[":
				if popval != "]" {
					bracket = append(bracket, "NO")
				}
				break
			case "{":
				if popval != "}" {
					bracket = append(bracket, "NO")
				}
				break
			case "(":
				if popval != ")" {
					bracket = append(bracket, "NO")
				}
				break
			}
			bracket = bracket[:len(bracket)-1]
		}
	}
	return len(bracket) == 0
}

func main() {
	var question = "[{(){}}]()]"
	answer := balanceBracket(question)
	fmt.Println(answer)
}


### Day 12: 29th Mar 2021

**Today's Progress:**
i learned how to count valleys using javascript and golang.

**Thoughts:**
basically create alt, if up ++ if down -- and if up check if alt -1 then numval++s

**Link to Work:**

javascript

function countValley(ar){
    let a = ar.split("")
    let alt = 0
    num_valley = 0
    for (let x = 0;x <= a.length - 1;x++){
        if (ar[x] == "U"){
            if(alt == -1){
                num_valley++
            }
            alt++
        }
        else if (ar[x] == "D"){
            alt--
        }
        else{console.log("err")}
    }
    return num_valley
}

let question = "DDUUDDUUDDUU"
let answer = countValley(question)
console.log(answer)

golang

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

### Day 13: 30th Mar 2021

**Today's Progress:**
i learned how to solve hackerrank count sock pair javascript and golang.

**Thoughts:**
basically create map and for loop the array. If there's map has no key of color then set new key of color. If map has then delete that key and numPair++.
ATTENTION : Javascript's Map different than golang. Javascript's map is an instance of a class so that to set, delete, and has need to use method.

**Link to Work:**

javascript

function countPair(ar){
    let mapPair = new Map()
    let numPair = 0
    for (x of ar){
        if (mapPair.has(x)){
            let tes = mapPair.delete(x)
            numPair++
        }
        else{
            mapPair.set(x,x)
        }
    }
    return numPair
}

let question = [1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7]
let answer = countPair(question)
console.log(answer)

Golang

package main

import (
	"fmt"
)

func countPair(ar []int) int {
	mapPair := make(map[int]int)
	numPair := 0
	for _, x := range ar {
		_, found := mapPair[x]
		if found == false {
			mapPair[x] = x
		} else {
			delete(mapPair, x)
			numPair++
		}
	}
	return numPair
}

func main() {
	var question = []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7}
	answer := countPair(question)
	fmt.Println(answer)
}

### Day 14: 31th Mar 2021

**Today's Progress:**
i learned how to solve hackerrank highest luck javascript and golang.

**Thoughts:**
basically sort multidim and for loop so that we can get important contest's luck value order after that if m is lower than k and value is important then += luckVal and m++.
The thing is if in javascript we just pass function to sort parameter, in golang we need to create type and create func of Len, Less, and Swap to sort

**Link to Work:**

javascript

function highestLuck(ar,k){
    let sortedArray = []
    let maxPoints = 0
    let m = 0
    ar.sort((a,b)=>{
        return a[0] - b[0]
    })
    sortedArray = ar.sort()
    console.log(sortedArray)
    for (let x = sortedArray.length-1;x>=0;x--){
        if (sortedArray[x][1] != 0 && m<k){
            maxPoints+=ar[x][0]
            m++
        }
        else if(sortedArray[x][1] == 0){
            maxPoints+=ar[x][0]
        }
    }
    return maxPoints
}

let question = [[1, 1], [2, 1], [3, 0], [4, 1], [5, 0], [6, 1]]
let answer = highestLuck(question,2)
console.log(answer)

golang

package main

import (
	"fmt"
	"sort"
)

type multiDim [][]int

func (a multiDim) Len() int {
	return len(a)
}

func (a multiDim) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a multiDim) Less(i, j int) bool {
	return a[j][0] < a[i][0]
}

func highestLuck(ar [][]int, k int) int {
	var sortedArray multiDim
	sortedArray = multiDim(ar)
	maxPoints := 0
	m := 0
	sort.Sort(multiDim(sortedArray))
	fmt.Println(sortedArray)
	for x := 0; x <= len(sortedArray)-1; x++ {
		if sortedArray[x][1] != 0 && m < k {
			m++
			maxPoints += sortedArray[x][0]
		} else if sortedArray[x][1] == 0 {
			maxPoints += sortedArray[x][0]
		}
	}
	return maxPoints
}

func main() {
	var question = [][]int{{1, 1}, {2, 1}, {3, 0}, {4, 1}, {5, 0}, {6, 1}}
	answer := highestLuck(question, 2)
	fmt.Println(answer)
}

### Day 15: 1st Apr 2021

**Today's Progress:**
i learned how to solve hackerrank find minimal absolute dif javascript and golang.

**Thoughts:**
basically double loop the array and if i==j continue, else check math.abs(ar[i]-ar[j]) if that lower than recently minimal value than assign it

**Link to Work:**

javascript

function lowestAbsDif(ar){
    let value = Number.POSITIVE_INFINITY
    for (let i = 0;i<=ar.length-1;i++){
        for(let j = 0;j<=ar.length-1;j++){
            if (j == i){
                continue
            }
            else{
                let dif = Math.abs(ar[i]-ar[j])
                if (dif < value){
                    value = dif
                }
            }
        }
    }
    return value
}

let question = [8,19,3]
let answer = lowestAbsDif(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func lowestAbsDif(ar []int) float64 {
	var min float64 = math.MaxFloat64
	for i := 0; i <= len(ar)-1; i++ {
		for j := 0; j <= len(ar)-1; j++ {
			if j == i {
				continue
			} else {
				dif := math.Abs(float64(ar[j]) - float64(ar[i]))
				if dif < min {
					min = dif
				}
			}
		}
	}
	return min
}

func main() {
	var question = []int{10, 1, 200}
	answer := lowestAbsDif(question)
	fmt.Println(answer)
}


### Day 16: 2nd Apr 2021

**Today's Progress:**
i learned how to solve hackerrank alternate character in javascript and golang.

**Thoughts:**
just for loop the splitted string then check if index x equal x+1 then continue, if x equal len - 1 or the last index append/push to stack and break the loop and else append/push to stack. Finally return stack.Join

**Link to Work:**

javascript

function alternateChar(s){
    let arrChar = s.split("")
    let charStack = []
    for (let x = 0;x <= arrChar.length-1;x++){
        if (x == arrChar.length-1){
            charStack.push(arrChar[x])
            break
        }
        else if (arrChar[x] == arrChar[x+1] && x+1 < arrChar.length){
            continue
        }
        else{
            charStack.push(arrChar[x])
        }
    }
    return charStack.join("")
}

let question = "AAABBBCCCDDDEEE"
let answer = alternateChar(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"strings"
)

func alternateChar(s string) string {
	arChar := strings.Split(s, "")
	var charStack []string
	for x, _ := range arChar {
		if x == len(arChar)-1 {
			charStack = append(charStack, arChar[x])
			break
		} else if arChar[x] == arChar[x+1] && x < len(arChar) {
			continue
		} else {
			charStack = append(charStack, arChar[x])
		}
	}

	return strings.Join(charStack, "")
}

func main() {
	var question string = "AAABBBCCCDDDEEE"
	answer := alternateChar(question)
	fmt.Println(answer)
}

### Day 17: 3rd Apr 2021

**Today's Progress:**
i learned how to solve hackerrank check analgam in javascript and golang.

**Thoughts:**
just check length, if not then false. if same convert string to []bye / UTF / ASCII then for loop and sum them and finally compare them
**Link to Work:**

javascript

function isAnagram(a,b){
    if (a.length != b.length){
        return false
    }
    let firstSum = 0
    let secondSum = 0
    for (let x=0;x<a.length;x++){
        firstSum+=a.charCodeAt(x)
        secondSum+=b.charCodeAt(x)
    }
    return firstSum == secondSum
}

let question = "PEOPLE"
let answer = isAnagram(question,"EPPLO")
console.log(answer)

golang

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

### Day 19: 5th Apr 2021

**Today's Progress:**
1) i learned how to solve hackerrank find toys in javascript and golang.
2) i learned how to solve hackerrank left rotation in javascript and golang.
3)

**Thoughts:**
1) similar to the highest luck challenge
2) in golang we just slice k: and then append :k... but in javascript we double for loop first push from index k to end then from index 0 to k
3)

**Link to Work:**

javascript

function toys(ar,k){
    ar.sort((a,b)=>{return b[0]-a[0]})
    let sortedToys = ar.sort()
    let recommendedToys = []
    let inc = 0
    for (let x=0;x<sortedToys.length && inc < k;x++){
        inc+=sortedToys[x][0]
        if (inc > k){
            break
        }
        recommendedToys.push(sortedToys[x][1])
    }
    return recommendedToys
}

let question = [[9,1],[3,2],[1,3],[5,4]]
let answer = toys(question,10)
console.log(answer)

function leftRotation(ar,k){
    if (k > ar.length - 1){
        return ar
    }
    let newAr = []
    for(let x=k;x<ar.length;x++){
        newAr.push(ar[x])
    }
    for(let x=0;x<k;x++){
        newAr.push(ar[x])
    }
    return newAr
}

let question = [1,2,3,4,5]
let answer = leftRotation(question,2)
console.log(answer)

golang

package main

import (
	"fmt"
	"sort"
)

type toys [][]int

func (t toys) Len() int {
	return len(t)
}

func (t toys) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t toys) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func findToys(ar [][]int, k int) []int {
	var sortedToys toys = toys(ar)
	sort.Sort(sortedToys)
	var recToys []int
	m := 0
	for x := 0; x < len(sortedToys) && m < k; x++ {
		m += sortedToys[x][0]
		if m > k {
			break
		}
		recToys = append(recToys, sortedToys[x][1])
	}
	return recToys
}

func main() {
	var question = [][]int{{9, 1}, {3, 2}, {1, 3}, {5, 4}}
	answer := findToys(question, 10)
	fmt.Println(answer)
}

package main

import (
	"fmt"
)

func leftRotation(ar []int, k int) []int {
	if k-1 > len(ar) {
		return ar
	}
	var newAr []int
	newAr = ar[k:]
	newAr = append(newAr, ar[:k]...)
	return newAr
}

func main() {
	var question = []int{1, 2, 3, 4, 5}
	answer := leftRotation(question, 2)
	fmt.Println(answer)
}







### Day 20: 6th Apr 2021

**Today's Progress:**
i learned how to solve hackerrank sum hourglass in javascript and golang.

**Thoughts:**
basically double while loop and sum the hourglass index relative

**Link to Work:**

javascript

function sumHourGlass(ar){
    let n = ar.length
    let sum = 0
    let i = 0
    while (i + 2 < n){
        let j = 0
        while (j + 2 < n){
            if (j + 2 >= n){
                break
            }
            let hourGlass = ar[i][j] + ar[i][j+1] + ar[i][j+2] + ar[i+2][j] + ar[i+2][j+1] + ar[i+2][j+2] + ar[i+1][j+1]
            sum += hourGlass
            j++
        }
        i++
    }
    return sum
}

let question = [[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1]]
let answer = sumHourGlass(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func sumHourGlass(ar [][]int) int {
	n := len(ar)
	i := 0
	sum := 0
	for {
		j := 0
		if i+2 >= n {
			break
		}
		for {
			if j+2 >= n {
				break
			}
			hourGlass := ar[i][j] + ar[i][j+1] + ar[i][j+2] + ar[i+1][j] + ar[i+1][j+1] + ar[i+1][j+2] + ar[i+1][j+1]
			sum += hourGlass
			j++
		}
		i++
	}
	return sum
}

func main() {
	var question = [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}
	answer := sumHourGlass(question)
	fmt.Println(answer)
}

### Day 21: 7th Apr 2021

**Today's Progress:**
i learned how to solve hackerrank two strings in javascript and golang.

**Thoughts:**
basically for loop and check contains for every word in argument 1

**Link to Work:**

javascript

function twoStrings(a1,a2){
    for (let x=0;x < a1.length; x++){
        if (a2.includes(a1[x])){
            return true
        }
    }
    return false
}

let question = "Hello"
let answer = twoStrings(question,"z")
console.log(answer)

golang

package main

import (
	"fmt"
	"strings"
)

func twoStrings(a1, a2 string) bool {
	for i, _ := range a1 {
		if strings.Contains(a2, string(a1[i])) == true {
			return true
		}
	}
	return false
}

func main() {
	var question string = "hello"
	answer := twoStrings(question, "zzz")
	fmt.Println(answer)
}

### Day 22: 8th Apr 2021

**Today's Progress:**
i learned how to solve leetcode number of island in javascript and golang.

**Thoughts:**
basically for loop and check if it 1 or island then recursively change that chain of 1 into 0.
In Golang use pointer btw.

**Link to Work:**

javascript

function numIsland(arr){
    let islandNum = 0
    for (let i=0;i < arr.length;i++){
        for(let j=0;j < arr[i].length;j++){
            if (arr[i][j] == 1){
                islandNum++
                findAndDestroy(arr,i,j)
            }
        }
    }
    return islandNum
}

function findAndDestroy(arr,i,j){
    if (i < 0 || i > arr.length-1 || j < 0 || j > arr[i].length - 1 || arr[i][j] == 0){
        return
    }
    arr[i][j] = 0
    findAndDestroy(arr,i+1,j)
    findAndDestroy(arr,i-1,j)
    findAndDestroy(arr,i,j+1)
    findAndDestroy(arr,i,j-1)
}

let question = [[1,1,1,1],[1,1,0,0],[1,0,0,0],[0,0,1,1],[1,0,0,0]]
let answer = numIsland(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func numIsland(arr [][]int) int {
	islandNum := 0
	for i, _ := range arr {
		for j, _ := range arr[i] {
			if arr[i][j] == 1 {
				islandNum++
				findAndDestroy(&arr, i, j)
			}
		}
	}
	return islandNum
}

func findAndDestroy(arr *[][]int, i int, j int) {
	if i < 0 || i > len(*arr)-1 || j < 0 || j > len((*arr)[i])-1 || (*arr)[i][j] == 0 {
		return
	}
	(*arr)[i][j] = 0
	findAndDestroy(arr, i+1, j)
	findAndDestroy(arr, i-1, j)
	findAndDestroy(arr, i, j+1)
	findAndDestroy(arr, i, j-1)
}

func main() {
	question := [][]int{{1, 1, 1, 1}, {1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 1, 1}, {1, 0, 0, 0}}
	answer := numIsland(question)
	fmt.Println(answer)
}

### Day 23: 9th Apr 2021

**Today's Progress:**
i learned how to solve leetcode pascal triangle in javascript and golang.

**Thoughts:**
basically double for loop, first loop is normal. Then create prevRow and NewRow. prevRow is index - 1 of triangle. Don't forget to continue if index = 0. After that second loop is begin from j = 1 to j < i. To be noticed that addition is always start from index 1 (because we'll add 1 num before) and j never exceeded i for example 1 3 3 1... number in the middle (3,3) is always lower than that index (3). Finally each loop is add previous item and current item [j-1]+ [j].

**Link to Work:**

javascript

function pascalTriangle(k){
    let triangle = []
    triangle.push([1])
    for (let i = 0;i < k;i++){
        if (i <= 0){
            continue
        }
        let prevRow = triangle[i-1]
        let newRow = [1]
        for(let j = 1; j < i; j++){
         newRow.push(prevRow[j-1]+prevRow[j])   
        }
        newRow.push(1)
        triangle.push(newRow)
    }
    return triangle
}

let question = 5
let answer = pascalTriangle(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func pascalTriangle(k int) [][]int {
	triangle := [][]int{}
	triangle = append(triangle, []int{1})
	for i := 0; i < k; i++ {
		if i <= 0 {
			continue
		} else {
			prevRow := triangle[i-1]
			newRow := []int{1}
			for j := 0; j < i-1; j++ {
				newRow = append(newRow, prevRow[j+1]+prevRow[j])
			}
			newRow = append(newRow, 1)
			triangle = append(triangle, newRow)
		}
	}
	return triangle
}

func main() {
	question := 5
	answer := pascalTriangle(question)
	fmt.Println(answer)
}


### Day 24: 10th Apr 2021

**Today's Progress:**
i learned how to solve leetcode find duplicate in javascript and golang.

**Thoughts:**
each item assign in map, if not found then set key, value if found then return true

**Link to Work:**

javascript

function findDuplicate(ar){
    let tmpMap = new Map()
    for (x of ar){
        if (tmpMap.get(x)){
            return true
        } else {
            tmpMap.set(x,x)
        }
    }
    return false
}

let question = [1,2,3,4,5,6,7,8,9]
let answer = findDuplicate(question)
console.log(answer)

golang

package main

import (
	"fmt"
)

func findDuplicate(ar []int) bool {
	tmpMap := make(map[int]int)
	for x := 0; x < len(ar); x++ {
		_, y := tmpMap[ar[x]]
		if y == false {
			tmpMap[ar[x]] = ar[x]
		} else {
			return true
		}
	}
	return false
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 9}
	answer := findDuplicate(question)
	fmt.Println(answer)
}

### Day 25: 11th Apr 2021

**Today's Progress:**
i learned how to solve leetcode two sums in javascript and golang.

**Thoughts:**
create two pointers based on index of array. One pointer is the first index and the other is the last index. Then using while loop sum them up and create two conditions. If sum < target then firspointer++ if sum > target then lastpointer--. else or == return the result.

**Link to Work:**

javascript

function twoSum(ar,target){
    let sum = ar[0] + ar[ar.length-1]
    let a_pointer = 0
    let b_pointer = ar.length - 1
    while (a_pointer <= b_pointer){
        sum = ar[a_pointer] + ar[b_pointer]
        if (sum < target){
            a_pointer++
        }else if(sum > target){
            b_pointer--
        }else {
            return [ar[a_pointer],ar[b_pointer]]
        }
    }
}

let question = [1,2,3,4,5,6,7,8]
let answer = twoSum(question,7)
console.log(answer)

golang

package main

import (
	"fmt"
)

func twoSum(ar []int, target int) []int {
	a_pointer := 0
	b_pointer := len(ar) - 1
	sum := ar[a_pointer] + ar[b_pointer]
	for {
		sum = ar[a_pointer] + ar[b_pointer]
		if sum < target {
			a_pointer++
		} else if sum > target {
			b_pointer--
		} else {
			return []int{ar[a_pointer], ar[b_pointer]}
		}
	}
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8}
	answer := twoSum(question, 8)
	fmt.Println(answer)
}

### Day 26: 12th Apr 2021

**Today's Progress:**
i learned how to solve leetcode palindrome number in javascript and golang.

**Thoughts:**
cast int to string (stringify), split that, for loop (reversed) and append to newAr... return newAr.join == stringify

**Link to Work:**

javascript

function palindromeNumber(a){
    let stringify = a.toString()
    let charArr = stringify.split("")
    let reversedArr = []
    for (let x=charArr.length-1;x>=0;x--){
        reversedArr.push(charArr[x])
    }
    return reversedArr.join("") == a
}

let question = 11111111
let answer = palindromeNumber(question)
console.log(answer)

golang

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


### Day 27: 13th Apr 2021

**Today's Progress:**
i learned how to solve binary search in javascript and golang.

**Thoughts:**
while loop, create first index and last index, and find middle index. if value of middle index is bigger than target than last index equal middle index and vice versa. Return index if middle == target. Don't forget to check if last index / first index == k

**Link to Work:**

javascript

function BinarySearch(ar,k){
    first_i = 0
    last_i = ar.length - 1
    middle_i = Math.ceil((first_i+last_i)/2)
    while (first_i <= last_i){
        if(ar[middle_i] < k){
            if (ar[last_i] == k){
                return last_i
            }
            first_i = middle_i
            middle_i = Math.ceil((first_i+last_i)/2)
        }else if(ar[middle_i] > k){
            if (ar[first_i] == k){
                return first_i
            }
            last_i = middle_i
            middle_i = Math.ceil((first_i+last_i)/2)
        }else if (ar[middle_i] == k){
            return middle_i
        }
    }
}

let question = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14]
let answer = BinarySearch(question,7)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func binarySearch(ar []int, k int) int {
	first_i := 0
	last_i := len(ar) - 1
	middle_i := math.Ceil((float64(first_i) + float64(last_i)) / 2)

	for {
		if ar[int(middle_i)] < k {
			if ar[last_i] == k {
				return last_i
			}
			first_i = int(middle_i)
			middle_i = math.Ceil((float64(first_i) + float64(last_i)) / 2)
		} else if ar[int(middle_i)] > k {
			if ar[first_i] == k {
				return first_i
			}
			last_i = int(middle_i)
			middle_i = math.Ceil((float64(first_i) + float64(last_i)) / 2)
		} else if ar[int(middle_i)] == k {
			return int(middle_i)
		}
	}
}

func main() {
	question := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	answer := binarySearch(question, 7)
	fmt.Println(answer)
}

### Day 28: 14th Apr 2021

**Today's Progress:**
i learned how to solve key and rooms challenge in javascript and golang.

**Thoughts:**
bikin 2 map (seen dan rooms), 1 stack (key).
awalnya for loop dulu masukin semua item di array ke map rooms.
push key dengan index 0 karena pertama kali akses ruangan index 0.
Kemudian pop si key sehingga didapat current key.
Setelah itu for loop setiap item di map menggunakan current key.
Pada setiap perulangan, cek rooms[curkey][i] apakah ada di map seen atau tidak.
Jika, tidak maka seen[rooms[curkey][i]] = true dan masukkan keynya ke stack.
key.push(rooms[curkey][i])
return len(seen) == len(array)

**Link to Work:**

javascript

function keyAndRoom(ar){
    let rooms = new Map()
    let seen = new Map()
    let keys = []
    seen.set(0,true)
    keys.push(0)
    for (let x=0;x < ar.length;x++){
        rooms.set(x,ar[x])
    }
    while (keys.length != 0){
        let curKey = keys.pop()
        
        for(let x = 0;x < rooms.get(curKey).length;x++){
            if (!seen.get(rooms.get(curKey)[x])){
                seen.set(rooms.get(curKey)[x],true)
                keys.push(rooms.get(curKey)[x])
            }
        }
    }
    return seen.size == ar.length
}

let question = [[1],[2],[3],[4],[5],[6],[2]]
let answer = keyAndRoom(question,7)
console.log(answer)

golang

package main

import (
	"fmt"
)

func keyandRoom(ar [][]int) bool {
	keys := []int{}
	rooms := make(map[int][]int)
	seen := make(map[int]bool)

	for x := 0; x < len(ar); x++ {
		rooms[x] = ar[x]
	}

	keys = append(keys, 0)
	seen[0] = true

	for {
		if len(keys) <= 0 {
			break
		}
		curKey := keys[len(keys)-1]
		keys = keys[:len(keys)-1]
		for i := 0; i < len(rooms[curKey]); i++ {
			if seen[rooms[curKey][i]] != true {
				seen[rooms[curKey][i]] = true
				keys = append(keys, rooms[curKey][i])
			}
		}
	}

	return len(seen) == len(ar)
}

func main() {
	question := [][]int{{1}, {2}, {3}, {4}, {5}, {6}}
	answer := keyandRoom(question)
	fmt.Println(answer)
}

### Day 29: 15th Apr 2021

**Today's Progress:**
i learned how to solve container with most water challenge in javascript and golang.

**Thoughts:**
create tmp = 0.
do double for loop. and then check if the heights (ar[i] < ar[j] or vice versa).
then check if size = math.abs(i - j) * ar[i] //or ar[j] is greater than tmp. If it is then
tmp = size

**Link to Work:**

javascript

function containerWithMostWater(ar){
    let tmpSize = 0
    for(let i = 0;i < ar.length;i++){
        for(let j = 0;j < ar.length;j++){
            if (j == i){
                continue
            }
            else {
                if (ar[i] < ar[j]){
                    let size = Math.abs((i-j)) * ar[i]
                    if (size > tmpSize){
                        tmpSize = size
                    }
                }else{
                    let size = Math.abs((i-j)) * ar[j]
                    if (size > tmpSize){
                        tmpSize = size
                    }
                }
            }
        }
    }
    return tmpSize
}

let question = [1,99999,2222,3,3333,5,7,90]
let answer = containerWithMostWater(question)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func containerWithMostWater(ar []int) float64 {
	tmp := 0.0
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[i] < ar[j] {
				size := math.Abs(float64(i)-float64(j)) * float64(ar[i])
				if size > tmp {
					tmp = size
				}
			} else {
				size := math.Abs(float64(i)-float64(j)) * float64(ar[j])
				if size > tmp {
					tmp = size
				}
			}
		}
	}
	return tmp
}

func main() {
	question := []int{1, 99999, 2222, 3, 3333, 5, 7, 90}
	answer := containerWithMostWater(question)
	fmt.Println(answer)
}

### Day 30: 16th Apr 2021

**Today's Progress:**
i learned how to solve leet code 3 sums challenge in javascript and golang.

**Thoughts:**
create for loop.
do two pointer technique for each loop with a_pointer = i+1 and b_pointer = len(ar)-1.
each loop calculate tmp = ar[i] + ar[a_pointer] + b[b_pointer].
check if |tmp - target| < |result - target|.
if it is, change result = tmp

**Link to Work:**

javascript

function threeSum(ar,k){
    let res = Number.POSITIVE_INFINITY
    let resAr = []
    for (let i = 0;i < ar.length;i++){
        let a_pointer = i+1
        let b_pointer = ar.length - 1
        while (a_pointer < b_pointer){
            let tmp = ar[a_pointer] + ar[b_pointer] + ar[i]
            if (ar[a_pointer] + ar[b_pointer] > k){
                b_pointer--
            }else{
                a_pointer++
            }
            if (Math.abs(tmp - k) < Math.abs(res - k)){
                res = tmp
                resAr = [ar[a_pointer],ar[b_pointer],ar[i]]
            }
        }
    }
	return resAr
}

let question = [1,2,3,4,5,6,7,8,9,10,11]
let answer = threeSum(question,19)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func threeSum(ar []int, k int) []int {
	res := math.MaxInt64
	resAr := []int{}
	for i := 0; i < len(ar); i++ {
		a_pointer := 0
		b_pointer := len(ar) - 1
		for {
			tmp := ar[i] + ar[a_pointer] + ar[b_pointer]
			if a_pointer >= b_pointer {
				break
			}
			if ar[a_pointer]+ar[b_pointer] > k {
				b_pointer--
			} else {
				a_pointer++
			}
			if math.Abs(float64(tmp)-float64(k)) < math.Abs(float64(res)-float64(k)) {
				res = tmp
				resAr = []int{ar[i], ar[a_pointer], ar[b_pointer]}
			}
		}
	}
	return resAr
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	answer := threeSum(question, 19)
	fmt.Println(answer)
}

### Day 31: 17th Apr 2021

**Today's Progress:**
i learned how to solve leet code hands of straight challenge in javascript and golang.

**Thoughts:**
check whether len(ar) % k != 0.
sort the array. 
double for loop.
first loop this conditions i = 0;i < len(ar);i+=totalGroup.
second loop this conditions j = i; j <= i + totalGroup - 1;j++.
each second loop append ar[j] to tmpAr.
each firs loop append tmpAr to resultAr.
return resultAr.

**Link to Work:**

javascript

function handOfStraight(ar,k){
    if (ar.length % k != 0){
        return []
    }else{
        ar.sort((a,b)=>{return a-b})
        let groupCount = ar.length / k
        let grouped = []
        for(let i = 0;i < ar.length;i+=groupCount){
            let tmp = []
            for(let j=i;j<=i+groupCount-1;j++){
                tmp.push(ar[j])
            }
            grouped.push(tmp)
        }
        return grouped
    }
}

let question = [12,11,10,9,8,7,6,5,4,3,2,1]
let answer = handOfStraight(question,4)
console.log(answer)

golang

package main

import (
	"fmt"
	"sort"
)

type cards []int

func (c cards) Len() int {
	return len(c)
}

func (c cards) Swap(a, b int) {
	c[a], c[b] = c[b], c[a]
}

func (c cards) Less(a, b int) bool {
	return c[a] < c[b]
}

func handOfStraight(ar []int, k int) [][]int {
	cards := cards(ar)
	sort.Sort(cards)
	groupModulus := len(ar) % k
	if groupModulus != 0 {
		return [][]int{}
	} else {
		var groupResult [][]int
		totalGroup := len(ar) / k
		for i := 0; i < len(ar); i += totalGroup {
			var grouped []int
			for j := i; j <= i+totalGroup-1; j++ {
				grouped = append(grouped, ar[j])
			}
			groupResult = append(groupResult, grouped)
		}
		return groupResult
	}
}

func main() {
	question := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	answer := handOfStraight(question, 4)
	fmt.Println(answer)
}

### Day 32: 18th Apr 2021

**Today's Progress:**
i learned how to solve leet code fastest method api call challenge in javascript and golang.

**Thoughts:**
Similar to binary search.
While loop a_pointer < b_pointer.
m_index = floor 1/2 * (a_index + b_index).
m_pointer = ar[m_index].
if apiCall(m_pointer) == true, then a_index = m_index, a_pointer = m_pointer.
else b_index = m_index + 1 ####+ 1 is required so that it breaks the loop
then return m_pointer

**Link to Work:**

javascript

function firstBadCode(ar,apiCall){
    let a_index = 0
    let b_index = ar.length - 1
    let mid_index = Math.floor((a_index+b_index)*0.5)
    let a_pointer = ar[a_index]
    let mid_pointer = ar[mid_index]
    let b_pointer = ar[b_index]
    let step = 0
    while (a_pointer < b_pointer){
        mid_index = Math.floor((a_index+b_index)*0.5)
        mid_pointer = ar[mid_index]
        let result = apiCall(mid_pointer)
        if (result == true){
            step++
            b_index = mid_index
            b_pointer = ar[b_index]
        }
        else{
            step++
            a_index = mid_index + 1
            a_pointer = ar[a_index]
        }
    }
    return `step: ${step}, first bad: ${ar[mid_pointer]}`
}

function ApiCall(version){
    if (version >= 17){
        return true
    } else{
        return false
    }
}

let question = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24]
let answer = firstBadCode(question, ApiCall)
console.log(answer)

golang

package main

import (
	"fmt"
	"math"
)

func firstBadCode(ar []int, apiCall func(s int) bool) string {
	a_index := 0
	b_index := len(ar) - 1
	a_pointer := ar[a_index]
	b_pointer := ar[b_index]
	step := 0
	m_index := int(math.Floor((float64(a_index) + float64(b_index)) * 0.5))
	m_pointer := ar[m_index]
	for {
		m_index = int(math.Floor((float64(a_index) + float64(b_index)) * 0.5))
		m_pointer = ar[m_index]
		if a_pointer >= b_pointer {
			break
		}
		if apiCall(m_pointer) == true {
			step++
			b_index = m_index
			b_pointer = ar[b_index]
		} else {
			step++
			a_index = m_index + 1
			a_pointer = ar[a_index]
		}
	}
	return fmt.Sprintf("step: %v firstbad : %v", step, m_pointer)
}

func apiCall(s int) bool {
	if s >= 17 {
		return true
	} else {
		return false
	}
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	answer := firstBadCode(question, apiCall)
	fmt.Println(answer)
}

### Day 32: 18th Apr 2021

**Today's Progress:**
i learned how to solve leet code return to origin challenge in javascript and golang.

**Thoughts:**
split the string.
for loop the array of string.
append to array for U, D, L, R.
return len(U) == len(D) && len(R) == len(L)

**Link to Work:**

javascript

function backOrigin(path){
    let pathAr = path.split("")
    let U = []
    let D = []
    let R = []
    let L = []
    for (let x=0;x<pathAr.length;x++){
        switch(pathAr[x]){
            case "U":
                U.push("U")
                break
            case "D":
                D.push("D")
                break
            case "R":
                R.push("R")
                break
            case "L":
                L.push("L")
                break
            default:
                return false
        }
    }

    return U.length == D.length && R.length == L.length
}

let question = "UUULLLDDDRRR"
let answer = backOrigin(question)
console.log(answer)

golang

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
