package main

import "testing"

var inputs = []int{
	3,
	4,
	9,
	58,
}

var expects = []string{
	"III",
	"IV",
	"IX",
	"LVIII",
}

func TestIntToRoman(t *testing.T) {
	for n, input := range inputs {
		if ans := intToRoman(input); ans != expects[n] {
			t.Errorf("rotate, input: %v, expected %s, but got: %v", input, expects[n], ans)
		}
	}
}

