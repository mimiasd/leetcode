package main

func missingNumber(nums []int) int {
	sumAll := (len(nums)+1)*(len(nums)+1+0)/ 2 
	sum := 0
	for _, i := range nums {
		sum += i
	}
	
	return sumAll - sum
}
