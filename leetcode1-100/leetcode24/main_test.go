package main

import "testing"

var inputs1 = [][]int{
	[]int{2,7,11,15},
	[]int{3,2,4},
	[]int{3,3},
}

var inputs2 = []int{
	9,
	6,
	6,
}

var expects = [][]int{
	[]int{0, 1},
	[]int{1, 2},
	[]int{0, 1},
}

func TestTwoSum(t *testing.T) {
	for n, input := range inputs1 {
		if ans := twoSum(input, inputs2[n]); !equalArray(ans, expects[n]) {
			t.Errorf("rotate, input1: %v, input2: %v, expected %d, but got: %v", input, inputs2[n], expects[n], ans)
		}
	}
}

func equalArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	for n, v := range a {
		if b[n] != v {
			return false
		}
	}

	return true
}
