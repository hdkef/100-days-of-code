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