package solution

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	prev1, prev2 := 1, 1

	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}

	return prev1
}
