package main

func removeDuplicates(nums []int) int {
	// 异常情况处理
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	// 正常情况处理
	v := 0 // 记录正在处理的值
	pos := 1 // 记录要被更改的位置

	for n, num := range nums {
		if n == 0 {
			v = num
		} else {
			if v != num {
				v = num
				nums[pos] = num
				pos += 1
			}
		}
	}

	return pos
}
