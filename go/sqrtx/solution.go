package solution

func mySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2
		mid_square := mid * mid
		if mid_square == x {
			return mid
		} else if mid_square < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}
