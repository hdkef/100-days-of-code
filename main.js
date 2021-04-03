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