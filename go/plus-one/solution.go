package solution

func plusOne(digits []int) []int {
	carry := 0
	length := len(digits) - 1

	digits[length] += 1
	for i := length; i >= 0; i-- {
		sum := digits[i] + carry
		carry, digits[i] = sum/10, sum%10
	}

	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
