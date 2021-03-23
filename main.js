function findBig(a){
    let biggestInt = 0
    try{
        for (let x = 0;x<=a.length-1;x++){
            if (a[x] > biggestInt){
                biggestInt = a[x]
            }
        }
        return biggestInt
    }catch(err){
        console.log(err)
    }
}

let question = [2,1,2,34,787,99999]
let answer = findBig(question)
console.log(answer)