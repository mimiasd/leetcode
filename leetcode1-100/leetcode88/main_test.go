package main

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	expect := []int{1, 2, 2, 3, 5, 6}
	if merge(nums1, 3, nums2, 3); !equalArray(nums1, expect) {
		t.Errorf("merge, expected: %v, but got: %v", expect, nums1)
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
