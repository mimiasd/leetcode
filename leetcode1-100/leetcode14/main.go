package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := 0
	for i := 0; i < len(strs); i++ {
		if i == 0 || minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}

	if minLen == 0 {
		return ""
	}

	var res strings.Builder
LOOP:
	for i := 0; i < minLen; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				break LOOP
			}
		}
		res.WriteByte(strs[0][i])
	}

	return res.String()
}
