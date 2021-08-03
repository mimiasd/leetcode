package main

import "testing"

var inputs = [][]int{
	[]int{1, 1, 1, 1, 3},
	[]int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4},
}

var expects = []int{
	3,
	9,
}

func TestRemoveDuplicates(t *testing.T) {
	for n, input := range inputs {
		if ans := removeDuplicates(input); ans != expects[n] {
			t.Errorf("RemoveDuplicates, input: %v. expected: %d, but got: %d", input, expects[n], ans)
		}
	}
}
