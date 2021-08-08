package main

import "testing"

var inputs = []string{
	"abcabcbb",
	"bbbbb",
	"pwwkew",
	"",
}

var expects = []int{
	3,
	1,
	3,
	0,
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for n, input := range inputs {
		if ans := lengthOfLongestSubstring(input); ans != expects[n] {
			t.Errorf("lengthOfLongestSubstring, input: %v, expected %d, but got: %v", input, expects[n], ans)
		}
	}
}

