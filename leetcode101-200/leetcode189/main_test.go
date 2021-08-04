package main

import "testing"

var inputs1 = [][]int{
	[]int{1, 2, 3, 4, 5, 6, 7},
	[]int{-1, -100, 3, 99},
}

var inputs2 = []int{
	3,
	2,
}

var inputs3 = [][]int{
	[]int{1, 2, 3, 4, 5, 6, 7},
	[]int{-1, -100, 3, 99},
}

var expects = [][]int{
	[]int{5, 6, 7, 1, 2, 3, 4},
	[]int{3, 99, -1, -100},
}

func TestRotate(t *testing.T) {
	for n, input := range inputs1 {
		if rotate(input, inputs2[n]); !equalArray(input, expects[n]) {
			t.Errorf("rotate, input1: %v, input2: %v, expected %d", input, inputs2[n], expects[n])
		}
	}
}

func TestRotate1(t *testing.T) {
	for n, input := range inputs3 {
		if rotate1(input, inputs2[n]); !equalArray(input, expects[n]) {
			t.Errorf("rotate, input1: %v, input2: %v, expected %d", input, inputs2[n], expects[n])
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
