function threeSum(ar,k){
    let res = Number.POSITIVE_INFINITY
    let resAr = []
    for (let i = 0;i < ar.length;i++){
        let a_pointer = i+1
        let b_pointer = ar.length - 1
        while (a_pointer < b_pointer){
            let tmp = ar[a_pointer] + ar[b_pointer] + ar[i]
            if (ar[a_pointer] + ar[b_pointer] > k){
                b_pointer--
            }else{
                a_pointer++
            }
            if (Math.abs(tmp - k) < Math.abs(res - k)){
                res = tmp
                resAr = [ar[a_pointer],ar[b_pointer],ar[i]]
            }
        }
        return resAr
    }
    
}

let question = [1,2,3,4,5,6,7,8,9,10,11]
let answer = threeSum(question,19)
console.log(answer)