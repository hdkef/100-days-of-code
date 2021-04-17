function firstBadCode(ar,apiCall){
    let a_index = 0
    let b_index = ar.length - 1
    let mid_index = Math.floor((a_index+b_index)*0.5)
    let a_pointer = ar[a_index]
    let mid_pointer = ar[mid_index]
    let b_pointer = ar[b_index]
    let step = 0
    while (a_pointer < b_pointer){
        mid_index = Math.floor((a_index+b_index)*0.5)
        mid_pointer = ar[mid_index]
        let result = apiCall(mid_pointer)
        if (result == true){
            step++
            b_index = mid_index
            b_pointer = ar[b_index]
        }
        else{
            step++
            a_index = mid_index + 1
            a_pointer = ar[a_index]
        }
    }
    return `step: ${step}, first bad: ${ar[mid_pointer]}`
}

function ApiCall(version){
    if (version >= 17){
        return true
    } else{
        return false
    }
}

let question = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24]
let answer = firstBadCode(question, ApiCall)
console.log(answer)