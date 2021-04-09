function pascalTriangle(k){
    let triangle = []
    triangle.push([1])
    for (let i = 0;i < k;i++){
        if (i <= 0){
            continue
        }
        let prevRow = triangle[i-1]
        let newRow = [1]
        for(let j = 1; j < i; j++){
         newRow.push(prevRow[j-1]+prevRow[j])   
        }
        newRow.push(1)
        triangle.push(newRow)
    }
    return triangle
}

let question = 5
let answer = pascalTriangle(question)
console.log(answer)