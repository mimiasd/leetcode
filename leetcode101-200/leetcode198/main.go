package main

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}

	return dp[n]
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
