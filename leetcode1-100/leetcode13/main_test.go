package main

import "testing"


var inputs  = []string{
	"III",
	"IV",
	"IX",
	"LVIII",
}

var expects= []int{
	3,
	4,
	9,
	58,
}

func TestIntToRoman(t *testing.T) {
	for n, input := range inputs {
		if ans := romanToInt(input); ans != expects[n] {
			t.Errorf("romanToInt, input: %v, expected %v, but got: %v", input, expects[n], ans)
		}
	}
}

