function lowestAbsDif(ar){
    let value = Number.POSITIVE_INFINITY
    for (let i = 0;i<=ar.length-1;i++){
        for(let j = 0;j<=ar.length-1;j++){
            if (j == i){
                continue
            }
            else{
                let dif = Math.abs(ar[i]-ar[j])
                if (dif < value){
                    value = dif
                }
            }
        }
    }
    return value
}

let question = [8,19,3]
let answer = lowestAbsDif(question)
console.log(answer)