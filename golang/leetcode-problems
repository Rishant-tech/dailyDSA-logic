//majority num problem

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
