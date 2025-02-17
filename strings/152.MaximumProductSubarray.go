package main

import "fmt"

func main() {
	nums := []int{2, 3, -2, 4}
	ans := maxProduct(nums)
	fmt.Println(ans)
}

func maxProduct(nums []int) int {
	n := len(nums) // Size of the array

	pre, suff := 1, 1
	ans := nums[0]

	for i := 0; i < n; i++ {
		// Reset pre and suff if they are zero
		if pre == 0 {
			pre = 1
		}
		if suff == 0 {
			suff = 1
		}

		// Update the prefix and suffix product
		pre *= nums[i]
		suff *= nums[n-i-1]

		// Update the answer with the maximum value between current ans, pre, and suff
		ans = max(ans, max(pre, suff))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
