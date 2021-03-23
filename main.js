function confirmEnding(a,b){
    return a.replace(" ","").substr(-1,1).toUpperCase() == b.toUpperCase()
}

let question = "NOTHING BUT NOTHING"
let answer = confirmEnding(question,"g")
console.log(answer)