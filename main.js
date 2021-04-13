function BinarySearch(ar,k){
    first_i = 0
    last_i = ar.length - 1
    middle_i = Math.ceil((first_i+last_i)/2)
    while (first_i <= last_i){
        if(ar[middle_i] < k){
            if (ar[last_i] == k){
                return last_i
            }
            first_i = middle_i
            middle_i = Math.ceil((first_i+last_i)/2)
        }else if(ar[middle_i] > k){
            if (ar[first_i] == k){
                return first_i
            }
            last_i = middle_i
            middle_i = Math.ceil((first_i+last_i)/2)
        }else if (ar[middle_i] == k){
            return middle_i
        }
    }
}

let question = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14]
let answer = BinarySearch(question,7)
console.log(answer)