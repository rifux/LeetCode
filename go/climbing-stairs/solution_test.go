package solution

import "testing"

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 1}, // 1
		{2, 2}, // 1+1 2
		{3, 3}, // 1+1+1 1+2 2+1
		{4, 5}, // 1+1+1+1 1+2+1 2+1+1 1+1+2 2+2
		{5, 8}, // 1+1+1+1+1 1+1+1+2 1+1+2+1 1+2+1+1 2+1+1+1 1+2+2 2+1+2 1+2+2
	}

	for _, test := range tests {
		ans := climbStairs(test.input)

		if ans != test.want {
			t.Fatalf("%d: expected %d to equal %d", test.input, ans, test.want)
		}
	}
}
