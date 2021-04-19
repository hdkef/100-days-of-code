function backspaceString(s1,s2){
    let s1Ar = []
    let s2Ar = []
    for (let i=0;i<s1.length;i++){
        if (s1[i]=="#"){
            s1Ar.pop()
        }else {
            s1Ar.push(s1[i])
        }
    }
    for (let i=0;i<s2.length;i++){
        if (s2[i]=="#"){
            s2Ar.pop()
        }else {
            s2Ar.push(s1[i])
        }
    }

    return s1Ar.values == s2Ar.values
}

let question = "U#BC##d#"
let answer = backspaceString(question,"xyZ###q#")
console.log(answer)