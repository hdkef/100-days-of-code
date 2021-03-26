function findPath(a){
    let pLen = a.length
    let i = -1
    let path = []
    while (i < pLen){
        i+=2
        switch (a[i]){
            case 0:
                path.push(i)
                break
            case 1:
                i-=1
                if (a[i] == 0){
                    path.push(i)
                }
                break
        }
    }
    return path
}

let question = [0,1,0,0,1,0,0,1,0,0,1,0]
let answer = findPath(question)
console.log(answer)