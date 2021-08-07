package main

import "testing"

var inputs = [][]int{
	[]int{3, 4, 5, 1, 2},
	[]int{4, 5, 6, 7, 0, 1, 2},
}

var expects = []int{
	1,
	0,
}

func TestFindMin(t *testing.T) {
	for n, input := range inputs {
		if ans := findMin(input); ans != expects[n] {
			t.Errorf("findMin, input: %v, expected %d, but got: %d", input, expects[n], ans)
		}
	}
}

