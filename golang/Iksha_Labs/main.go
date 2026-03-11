package main

import "fmt"

// ═══════════════════════════════════════════════
// 1. TWO SUM
// Input: [2,7,11,15], target=9 → [0,1]
// ═══════════════════════════════════════════════
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value → index
	for i, num := range nums {
		if idx, ok := seen[target-num]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

// ═══════════════════════════════════════════════
// 2. CONTAINS DUPLICATE
// Input: [1,2,3,1] → true
// ═══════════════════════════════════════════════
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// ═══════════════════════════════════════════════
// 3. VALID ANAGRAM
// Input: "anagram", "nagaram" → true
// ═══════════════════════════════════════════════
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

// ═══════════════════════════════════════════════
// 4. VALID PARENTHESES
// Input: "()[]{}" → true | "([)]" → false
// ═══════════════════════════════════════════════
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// ═══════════════════════════════════════════════
// 5. REVERSE A STRING
// Input: "hello" → "olleh"
// ═══════════════════════════════════════════════
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ═══════════════════════════════════════════════
// 6. LONGEST SUBSTRING WITHOUT REPEATING CHARS
// Input: "abcabcbb" → 3 ("abc")
// ═══════════════════════════════════════════════
func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int) // char → last index
	maxLen, left := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := seen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		seen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// ═══════════════════════════════════════════════
// 7. MAXIMUM SUM SUBARRAY OF SIZE K
// Input: [1,3,2,6,4], k=3 → 12
// ═══════════════════════════════════════════════
func maxSumSubarray(nums []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

// ═══════════════════════════════════════════════
// 8. MAXIMUM SUBARRAY (Kadane's Algorithm)
// Input: [-2,1,-3,4,-1,2,1,-5,4] → 6
// ═══════════════════════════════════════════════
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// ═══════════════════════════════════════════════
// 9. BINARY SEARCH
// Input: [1,3,5,7,9], target=7 → 3
// ═══════════════════════════════════════════════
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// ═══════════════════════════════════════════════
// 10. BEST TIME TO BUY AND SELL STOCK
// Input: [7,1,5,3,6,4] → 5 (buy at 1, sell at 6)
// ═══════════════════════════════════════════════
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxPro := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxPro {
			maxPro = price - minPrice
		}
	}
	return maxPro
}

// ═══════════════════════════════════════════════
// 11. CLIMBING STAIRS
// Input: n=4 → 5 ways
// ═══════════════════════════════════════════════
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev1 = prev2
		prev2 = curr
	}
	return prev2
}

// ═══════════════════════════════════════════════
// LINKED LIST NODE
// ═══════════════════════════════════════════════
type ListNode struct {
	Val  int
	Next *ListNode
}

// ═══════════════════════════════════════════════
// 12. REVERSE LINKED LIST
// Input: 1→2→3→4→5 → 5→4→3→2→1
// ═══════════════════════════════════════════════
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// ═══════════════════════════════════════════════
// 13. LINKED LIST CYCLE DETECTION (Floyd's)
// Input: 1→2→3→4→2(cycle) → true
// ═══════════════════════════════════════════════
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// ═══════════════════════════════════════════════
// 14. MERGE TWO SORTED LISTS
// Input: 1→2→4, 1→3→4 → 1→1→2→3→4→4
// ═══════════════════════════════════════════════
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// ═══════════════════════════════════════════════
// BINARY TREE NODE
// ═══════════════════════════════════════════════
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ═══════════════════════════════════════════════
// 15. MAXIMUM DEPTH OF BINARY TREE
// ═══════════════════════════════════════════════
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

// ═══════════════════════════════════════════════
// 16. INVERT BINARY TREE
// ═══════════════════════════════════════════════
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// ═══════════════════════════════════════════════
// 17. SAME TREE
// ═══════════════════════════════════════════════
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// ═══════════════════════════════════════════════
// 18. WORKER POOL (Golang specific — very relevant)
// ═══════════════════════════════════════════════
// func workerPool() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
//
// 	// Start 5 workers
// 	for w := 1; w <= 5; w++ {
// 		go func(id int) {
// 			for job := range jobs {
// 				fmt.Printf("Worker %d processing job %d\n", id, job)
// 				results <- job * 2
// 			}
// 		}(w)
// 	}
//
// 	// Send 20 jobs
// 	for j := 1; j <= 20; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// }

// ═══════════════════════════════════════════════
// 19. RATE LIMITER (Token Bucket — very relevant)
// ═══════════════════════════════════════════════
// type RateLimiter struct {
// 	tokens   int
// 	maxTokens int
// 	mu       sync.Mutex
// }
//
// func (r *RateLimiter) Allow() bool {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	if r.tokens > 0 {
// 		r.tokens--
// 		return true
// 	}
// 	return false
// }

// ═══════════════════════════════════════════════
// QUICK REFERENCE — TIME COMPLEXITY
// ═══════════════════════════════════════════════
// Two Sum              → O(n) time, O(n) space
// Contains Duplicate   → O(n) time, O(n) space
// Valid Parentheses    → O(n) time, O(n) space
// Longest Substring    → O(n) time, O(n) space
// Binary Search        → O(log n) time, O(1) space
// Max Subarray         → O(n) time, O(1) space
// Reverse Linked List  → O(n) time, O(1) space
// Max Depth Tree       → O(n) time, O(h) space

func main() {
	// Test Two Sum
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]

	// Test Contains Duplicate
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // true

	// Test Valid Parentheses
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("([)]"))   // false

	// Test Reverse String
	fmt.Println(reverseString("hello")) // olleh

	// Test Longest Substring
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3

	// Test Binary Search
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 7)) // 3

	// Test Max Profit
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5

	// Test Climbing Stairs
	fmt.Println(climbStairs(4)) // 5
}
