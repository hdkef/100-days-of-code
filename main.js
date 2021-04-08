function numIsland(arr){
    let islandNum = 0
    for (let i=0;i < arr.length;i++){
        for(let j=0;j < arr[i].length;j++){
            if (arr[i][j] == 1){
                islandNum++
                findAndDestroy(arr,i,j)
            }
        }
    }
    return islandNum
}

function findAndDestroy(arr,i,j){
    if (i < 0 || i > arr.length-1 || j < 0 || j > arr[i].length - 1 || arr[i][j] == 0){
        return
    }
    arr[i][j] = 0
    findAndDestroy(arr,i+1,j)
    findAndDestroy(arr,i-1,j)
    findAndDestroy(arr,i,j+1)
    findAndDestroy(arr,i,j-1)
}

let question = [[1,1,1,1],[1,1,0,0],[1,0,0,0],[0,0,1,1],[1,0,0,0]]
let answer = numIsland(question)
console.log(answer)