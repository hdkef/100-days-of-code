function balanceBracket(ar){
    let a = ar.split("")
    let bracketStack = []
    for (let x = 0;x <= a.length - 1;x++){
        if (a[x] == '[' || a[x] == '{' || a[x] == '('){
            bracketStack.push(1)
        }
        else{
            if (bracketStack.length == 0){
                bracketStack.push(1)
                break
            }
            else{bracketStack.pop()}
        }
    }
    return bracketStack.length == 0
}

let question = "[{(){}}]"
let answer = balanceBracket(question)
console.log(answer)