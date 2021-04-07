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
