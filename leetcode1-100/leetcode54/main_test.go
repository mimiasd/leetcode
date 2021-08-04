package main

import "testing"

var input = [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
}

var expect = []int{1,2,3,6,9,8,7,4,5}

func TestSpiralOrder(t *testing.T) {
		if ans := spiralOrder(input); !equalArray(ans, expect) {
			t.Errorf("RemoveDuplicates, input: %v, expected: %d, but got: %d", input, expect, ans)
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
