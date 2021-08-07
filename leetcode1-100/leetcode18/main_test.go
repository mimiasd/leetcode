package main

import "testing"

var inputs = [][]int{
	[]int{1, 0, -1, 0, -2, 2},
	[]int{0},
}

var target int = 0

var expects = [][][]int{
	[][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
	[][]int{},
}

func TestFourSum(t *testing.T) {
	for n, input := range inputs {
		ans := fourSum(input, target)
		t.Logf("fourSum, input: %v, expected %d, got: %v", input, expects[n], ans)
	}
}
