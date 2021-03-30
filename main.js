function balanceBracket(ar){
    let a = ar.split("")
    let bracketStack = []
    for (let x = 0;x <= a.length - 1;x++){
        if (a[x] == '[' || a[x] == '{' || a[x] == '('){
            bracketStack.push(a[x])
        }
        else{
            if (bracketStack.length == 0){
                bracketStack = ["NO"]
                break
            }
            else{
                let popval = bracketStack.pop()
                console.log(popval)
                switch (a[x]){
                    case "}":
                        if (popval != "{"){
                            bracketStack = ["NO"]
                        }
                        break
                    case "]":
                        if (popval != "["){
                            bracketStack = ["NO"]
                        }
                        break
                    case ")":
                        if (popval != "("){
                            bracketStack = ["NO"]
                        }
                        break
                }
            }
        }
    }
    console.log(bracketStack)
    return bracketStack.length == 0
}

let question = "[{(){}}][]"
let answer = balanceBracket(question)
console.log(answer)