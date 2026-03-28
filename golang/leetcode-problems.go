/*majority num problem
with floor value
*/
//brute force - O(n2) - not accepted 

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

//optimize - with soring -- accepted 

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


// majority element will always exists >n/2, so center element will always the majority

func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

//moore algo - do not reset the frequency just increase if mj matches to current indexed element in array - O(n)

func majorityElement(nums []int) int {
    freq, mj := 0,0
    for i:=0; i<len(nums); i++{
        if freq ==0{
            mj = nums[i]
        }
        if (mj==nums[i]){
            freq +=1
        }else{
            freq-=1
        }
    }
    return mj
}


/*
Two Sum problem
*/

//brute force

func twoSum(nums []int, target int) []int {
    for i:=0 ; i<len(nums);i++{
        for j:=i+1; j<len(nums);j++{
            if nums[i]+nums[j] == target{
                return []int{i,j}
            }
        }
    }
    return nil
}
