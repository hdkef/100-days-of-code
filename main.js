function handOfStraight(ar,k){
    if (ar.length % k != 0){
        return []
    }else{
        ar.sort((a,b)=>{return a-b})
        let groupCount = ar.length / k
        let grouped = []
        for(let i = 0;i < ar.length;i+=groupCount){
            let tmp = []
            for(let j=i;j<=i+groupCount-1;j++){
                tmp.push(ar[j])
            }
            grouped.push(tmp)
        }
        return grouped
    }
}

let question = [12,11,10,9,8,7,6,5,4,3,2,1]
let answer = handOfStraight(question,4)
console.log(answer)