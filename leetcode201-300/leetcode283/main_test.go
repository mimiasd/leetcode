package main

import "testing"

var inputs = [][]int{
	[]int{0, 2, 2, 0},
	[]int{0, 1, 2, 2, 3, 0, 4, 2},
}

var expects = [][]int{
	[]int{2, 2, 0, 0},
	[]int{1, 2, 2, 3, 4, 2, 0, 0},
}

func TestRemoveDuplicates(t *testing.T) {
	for n, input := range inputs {
		if moveZeroes(input);  !equalArray(input, expects[n]) {
			t.Errorf("RemoveDuplicates, input: %v. expected: %d", input, expects[n])
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
