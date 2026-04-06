func lengthOfLongestSubstring(s string) int {

    //lets try it with brute force
    maxSize := 0 
    for i:=0; i<len(s);i++{
        //declare a map to keep track of seen elements
        seenMap := make(map[byte]bool)//, len(s))
        
        //two loops because we have to trace longest non repeating
        for j := i; j < len(s); j++ {
            if _,ok:= seenMap[s[j]];ok{
                break
            }
            seenMap[s[j]] = true

            if len(seenMap) > maxSize{
                maxSize = len(seenMap)
            }
        }
    }
    return maxSize
}

//sliding window concept 

// Template:
func fixedWindow(arr []int, k int) {
    windowSum := 0

   //create first window of size k, given
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    result := windowSum

   //now slice the window as per condition i.e, fixed or variable
    for i := k; i < len(arr); i++ {
        windowSum += arr[i]       // add new element (right)
        windowSum -= arr[i-k]     // remove first element (left)
        // result update 
        if windowSum > result {
            result = windowSum
        }
    }
}

func isValid(s string) bool {
    //declare stack
    stack := []rune{}

    matchMap := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i:=0; i<len(s);i++{
        char := rune(s[i])

        //match with rune type not with str type, otherwise type case error will be there
        if char == '(' || char == '{' || char == '['{
            stack = append(stack , char)
        }else{
            if len(stack) == 0{
                return false // as no open bracket seen so string is invalid
            }
            //extrack top elementa dn pop that from stack
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1] 

            //if not matched in matched map, then it is invalid
            if top != matchMap[char] {
                return false
            }


        }
    }
    return len(stack)==0
}

func isValid(s string) bool {

    //naive approach to replace the pairs with "" if there is a string pair found
    for strings.Contains(s,"()") || strings.Contains(s,"{}") || strings.Contains(s,"[]"){
        s = strings.ReplaceAll(s, "()", "")
        s = strings.ReplaceAll(s, "[]", "")
        s = strings.ReplaceAll(s, "{}", "")
    }
    return len(s) == 0
}

func longestCommonPrefix(strs []string) string {
    //return in case of empty string
    if len(strs) == 0{
        return ""
    }

    ans := ""

    //sort the string array and then compare first and last as strs b/w already sorted in order they contains common chars
    sort.Strings(strs)
    first := strs[0]
    last := strs[len(strs)-1]

    //there are 2 methods to traverse, first is to traverse with normal conditon and second one is to directly put all conditions and just increase i pointer, I will keep traversing with normal looping

    for i:=0; i<len(first); i++{
        //break the loop if index is greater then last string length - prevent out of bound
        if i>=len(last){
            break
        }
        if first[i]!=last[i]{
            break
        }
        //add up char to answer
        ans += string(first[i])
    }
    return ans
}

func longestCommonPrefix(strs []string) string {
    //return in case of empty string
    if len(strs) == 0{
        return ""
    }

    //lets try naive approach first
    ans := ""

    for i:=0; i<len(strs[0]);i++ {
        //take first str of array and itrate over that
        char := strs[0][i]
        for j:=1; j<len(strs);j++{
            if i>=len(strs[j]) || strs[j][i] != char{
                return ans
            }
        }
        ans += string(char)
    }
    return ans
}

func romanToInt(s string) int {
    romanMap := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000,
    }

    result := 0

    for i := 0; i < len(s); i++ {
        curr := romanMap[s[i]]

        if i+1 < len(s) && curr < romanMap[s[i+1]] {
            result -= curr  
        } else {
            result += curr 
        }
    }

    return result
}

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
