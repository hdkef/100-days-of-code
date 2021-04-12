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