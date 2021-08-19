package main

func isValid(s string) bool {
	if s == "" {
		return true
	}

	matchs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		v, ok := matchs[s[i]]
		if !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
		
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
