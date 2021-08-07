package main

import "testing"

var inputs = [][]int{
	[]int{1, 2, 3, 4},
	[]int{-1, 1, 0, -3, 3},
}

var expects = [][]int{
	[]int{24, 12, 8, 6},
	[]int{0, 0, 9, 0, 0},
}

func TestProductExceptSelf(t *testing.T) {
	for n, input := range inputs {
		if ans := productExceptSelf(input);  !equalArray(ans, expects[n]) {
			t.Errorf("RemoveDuplicates, input: %v. expected: %v, but got: %v", input, expects[n], ans)
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
