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