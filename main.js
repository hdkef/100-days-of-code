function countPair(ar){
    let mapPair = new Map()
    let numPair = 0
    for (x of ar){
        if (mapPair.has(x)){
            let tes = mapPair.delete(x)
            numPair++
        }
        else{
            mapPair.set(x,x)
        }
    }
    return numPair
}

let question = [1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7]
let answer = countPair(question)
console.log(answer)