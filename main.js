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