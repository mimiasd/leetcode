package main

func hammingWeight(num uint32) int {
	ones := 0
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return ones
}
