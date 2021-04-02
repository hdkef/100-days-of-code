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