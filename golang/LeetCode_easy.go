func isPalindrome(x int) bool {
    //one method can be to conver this into a string and the back traverse the string finally compare each other
    
    //if a number is negative then it can never be a palindrome so return at very first step
    if x<0{
        return false
    }
    reverseNum := 0

    //keep trace of orignal number
    orignalNum := x

    //traverse on for loop condition until the actual number is greater then 0
    for x >0 {
        //we have to keep 2 conditions running - find remainder and remove last digit
        reverseNum = reverseNum*10 + x%10
        x = x/10
    }
    //compare and return the number
    return reverseNum == orignalNum
}

func twoSum(nums []int, target int) []int {
    //method one can be - use, two loops and add each element and check with target, that will create an issue of high time complexity - O(n*n)

    //second can be to use maps and compare the values  - O(n)
    seenMap := make(map[int]int, len(nums)) //declare a int of int map to storekey value pair
    for i:=0; i<len(nums); i++{
        // we will apply simple math equation to find two nums lets say them a, b
        // eq => a+b = target and in reverse we can have target -b = a

        b := target - nums[i] // for instance b is to be find and nums[i] is our a

        //check if key - b exists in the map or not

        if key, ok := seenMap[b]; ok{
            // we found then return the indices of values that are making the target as a sum
            return []int{key, i}
        }else{
            //add the value to the map
            seenMap[nums[i]] = i
        }
    }
    return nil
}
