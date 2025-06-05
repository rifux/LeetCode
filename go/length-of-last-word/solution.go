package solution

func lengthOfLastWord(s string) (result int) {
	for r := len(s) - 1; r >= 0; r-- {
		if s[r] != ' ' {
			result++
		} else if result != 0 {
			return
		}
	}

	return
}
