function backOrigin(path){
    let pathAr = path.split("")
    let U = []
    let D = []
    let R = []
    let L = []
    for (let x=0;x<pathAr.length;x++){
        switch(pathAr[x]){
            case "U":
                U.push("U")
                break
            case "D":
                D.push("D")
                break
            case "R":
                R.push("R")
                break
            case "L":
                L.push("L")
                break
            default:
                return false
        }
    }

    return U.length == D.length && R.length == L.length
}

let question = "UUULLLDDDRRR"
let answer = backOrigin(question)
console.log(answer)