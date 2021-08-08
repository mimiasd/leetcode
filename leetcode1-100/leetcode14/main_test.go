package main

import "testing"

var inputs = [][]string{
	[]string{"flower", "flow", "flight"},
	[]string{"dog", "racecar", "car"},
	[]string{"", "a"},
	[]string{"ba", "a"},
}

var expects = []string{
	"fl",
	"",
	"",
	"",
	"",
}

func TestLongestCommonPrefix(t *testing.T) {
	for n, input := range inputs {
		if ans := longestCommonPrefix(input); ans != expects[n] {
			t.Errorf("longestCommonPrefix, input: %v, expected %s, but got: %v", input, expects[n], ans)
		}
	}
}

