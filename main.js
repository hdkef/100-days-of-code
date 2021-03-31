function highestLuck(ar,k){
    let sortedArray = []
    let maxPoints = 0
    let m = 0
    ar.sort((a,b)=>{
        return a[0] - b[0]
    })
    sortedArray = ar.sort()
    console.log(sortedArray)
    for (let x = sortedArray.length-1;x>=0;x--){
        if (sortedArray[x][1] != 0 && m<k){
            maxPoints+=ar[x][0]
            m++
        }
        else if(sortedArray[x][1] == 0){
            maxPoints+=ar[x][0]
        }
    }
    return maxPoints
}

let question = [[1, 1], [2, 1], [3, 0], [4, 1], [5, 0], [6, 1]]
let answer = highestLuck(question,2)
console.log(answer)