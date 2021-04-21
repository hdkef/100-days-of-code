function reveseOnlyLetter(s){
    letterOnly = []
    reversed = []
    for(let i=0;i<s.length;i++){
        if (s[i] != "-"){
            letterOnly.push(s[i])
        }
    }
    for(let j=0;j<s.length;j++){
        if(s[j] != "-"){
            reversed.push(letterOnly.pop())
        }else{
            reversed.push(s[j])
        }
    }
    return reversed.join("")
}

let question = "bismi-lah-irahman"
let answer = reveseOnlyLetter(question)
console.log(answer)