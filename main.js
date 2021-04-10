function findDuplicate(ar){
    let tmpMap = new Map()
    for (x of ar){
        if (tmpMap.get(x)){
            return true
        } else {
            tmpMap.set(x,x)
        }
    }
    return false
}

let question = [1,2,3,4,5,6,7,8,9]
let answer = findDuplicate(question)
console.log(answer)