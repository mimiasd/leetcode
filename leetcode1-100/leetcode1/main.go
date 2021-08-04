package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		ans := target - num 
		if _, ok := m[ans]; ok {
			return []int{m[ans], i}
		}
		m[num] = i
	}
	return nil
}
