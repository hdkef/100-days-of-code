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