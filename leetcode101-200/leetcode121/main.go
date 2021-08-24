package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res, mx := 0, prices[0]

	for i := 0; i < len(prices); i++ {
		res = max(res, prices[i]-mx)
		mx = min(mx, prices[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
