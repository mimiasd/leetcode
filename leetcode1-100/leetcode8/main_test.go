package main

import "testing"

var inputs = []string{
	"123",
	"-111",
	"+111",
	"1124a2",
	"   -42",
	"words and 987",
	"-91283472332",
}

var expects = []int{
	123,
	-111,
	111,
	1124,
	-42,
	0,
	-2147483648,
}

func TestMyAtoi(t *testing.T) {
	for n, input := range inputs {
		if ans := myAtoi(input);  ans != expects[n]{
			t.Errorf("rotate, input: %v, expected %d, but got: %v", input, expects[n], ans)
		}
	}
}

