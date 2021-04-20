function jewelsAndStones(j,s){
    let count = 0
    
    for(let i=0;i < s.length; i++) {
        for(let k=0;k < j.length;k++){
            if (s[i] == j[k]){
                count++
            }
        }
    }

    return count
}

let question = "abc"
let answer = jewelsAndStones(question,"aaabbcc")
console.log(answer)