function countValley(ar){
    let a = ar.split("")
    let alt = 0
    num_valley = 0
    for (let x = 0;x <= a.length - 1;x++){
        if (ar[x] == "U"){
            if(alt == -1){
                num_valley++
            }
            alt++
        }
        else if (ar[x] == "D"){
            alt--
        }
        else{console.log("err")}
    }
    return num_valley
}

let question = "DDUUDDUUDDUU"
let answer = countValley(question)
console.log(answer)