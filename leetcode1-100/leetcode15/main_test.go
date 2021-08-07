package main

import "testing"

var inputs = [][]int{
	[]int{-1, 0, 1, 2, -1, -4},
	[]int{},
	[]int{0},
}

var expects = [][][]int{
	[][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
	[][]int{},
	[][]int{},
}

func TestThreeSum(t *testing.T) {
	for n, input := range inputs {
		ans := threeSum(input)
		t.Logf("threeSum, input: %v, expected %d, got: %v", input, expects[n], ans)
	}
}
