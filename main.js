function diagonalDiff(m){
    var leftToRight = 0
    var rightToLeft = 0
    let i = 0
    while (i <= m.length - 1){
        leftToRight += m[i][i]
        console.log(m[i][i])
        let j = m.length - 1 - i
        rightToLeft += m[i][j]
        console.log(m[i][j])
        i++
        console.log(rightToLeft,leftToRight)
    }
    return Math.abs(leftToRight - rightToLeft)
}

let question = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]
let answer = diagonalDiff(question)
console.log(answer)