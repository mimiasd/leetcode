package main

import "testing"

var inputs1 = [][]int{
	[]int{3, 2, 2, 3},
	[]int{0, 1, 2, 2, 3, 0, 4, 2},
}

var inputs2 = []int{
	3,
	2,
}

var expects = []int{
	2,
	5,
}

func TestRemoveDuplicates(t *testing.T) {
	for n, input1 := range inputs1 {
		if ans := removeElement(input1, inputs2[n]); ans != expects[n] {
			t.Errorf("RemoveDuplicates, input1: %v, input2: %v. expected: %d, but got: %d", input1, inputs2[n], expects[n], ans)
		}
	}
}
