package main

func lengthOfLongestSubstring(s string) int {
	win := make(map[byte]int)
	n := len(s)
	ans := 0
	l, r := 0, 0

	for r < n {
		v := s[r]
		r++
		win[v]++
		for win[v] > 1 {
			win[s[l]]--
			l++
		}
		ans = max(ans, r - l)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
