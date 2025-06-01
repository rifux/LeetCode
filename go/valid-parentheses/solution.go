package validparentheses

func isValid(s string) bool {
	stack := []rune{}
	openers := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if opener, ok := openers[char]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != opener {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
