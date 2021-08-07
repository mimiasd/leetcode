package main

import "sort"


//  本题的难点在于如何去除重复解
func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)

	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n - 2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		j, k := i + 1, n -1
		for j < k {
			if nums[i] + nums[j] + nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				for k > 1 && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i] + nums[j] + nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
