package main

func rotate(nums []int, k int)  {
	n := len(nums)
	newNums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%n] = nums[i]	
	}

	copy(nums, newNums)
}

func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
