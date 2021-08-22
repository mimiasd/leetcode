package main

func maxSubArray(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	ans := nums[0]

	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		
		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
