/*majority num problem
with floor value
*/
//brute force - O(n2)

func majorityElement(nums []int) int {
    var mj int
	for _, i := range nums{
        freq := 0
		for _,j := range nums {
            if i==j{
                freq +=1
            }
        }
        if freq > len(nums)/2 {
            mj = i
        }
	}
    return mj
}

//optimize - with soring

func majorityElement(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    mj := nums[0]
    freq := 1
    if freq >n/2{
        return mj
    }
    for i:=1;i<n;i++{
        if nums[i] == nums[i-1]{
            freq+=1
        }else{
            freq=1
            mj=nums[i]
        }
        if freq>n/2{
            return mj
        }
    }
    return -1
}
