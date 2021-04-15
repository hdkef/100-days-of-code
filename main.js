function containerWithMostWater(ar){
    let tmpSize = 0
    for(let i = 0;i < ar.length;i++){
        for(let j = 0;j < ar.length;j++){
            if (j == i){
                continue
            }
            else {
                if (ar[i] < ar[j]){
                    let size = Math.abs((i-j)) * ar[i]
                    if (size > tmpSize){
                        tmpSize = size
                    }
                }else{
                    let size = Math.abs((i-j)) * ar[j]
                    if (size > tmpSize){
                        tmpSize = size
                    }
                }
            }
        }
    }
    return tmpSize
}

let question = [1,99999,2222,3,3333,5,7,90]
let answer = containerWithMostWater(question)
console.log(answer)