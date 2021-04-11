function twoSum(ar,target){
    let sum = ar[0] + ar[ar.length-1]
    let a_pointer = 0
    let b_pointer = ar.length - 1
    while (a_pointer <= b_pointer){
        sum = ar[a_pointer] + ar[b_pointer]
        if (sum < target){
            a_pointer++
        }else if(sum > target){
            b_pointer--
        }else {
            return [ar[a_pointer],ar[b_pointer]]
        }
    }
}

let question = [1,2,3,4,5,6,7,8]
let answer = twoSum(question,7)
console.log(answer)