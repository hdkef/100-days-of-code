function leftRotation(ar,k){
    if (k > ar.length - 1){
        return ar
    }
    let newAr = []
    for(let x=k;x<ar.length;x++){
        newAr.push(ar[x])
    }
    for(let x=0;x<k;x++){
        newAr.push(ar[x])
    }
    return newAr
}

let question = [1,2,3,4,5]
let answer = leftRotation(question,2)
console.log(answer)