package main

func removeDuplicates(nums []int) int {
	// 异常情况处理
	n := len(nums)
	if n <= 2 {
		return n
	}

	// 基本情况处理
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
