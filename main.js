function sumHourGlass(ar){
    let n = ar.length
    let sum = 0
    let i = 0
    while (i + 2 < n){
        let j = 0
        while (j + 2 < n){
            if (j + 2 >= n){
                break
            }
            let hourGlass = ar[i][j] + ar[i][j+1] + ar[i][j+2] + ar[i+2][j] + ar[i+2][j+1] + ar[i+2][j+2] + ar[i+1][j+1]
            sum += hourGlass
            j++
        }
        i++
    }
    return sum
}

let question = [[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1]]
let answer = sumHourGlass(question)
console.log(answer)