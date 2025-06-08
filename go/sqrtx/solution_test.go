package solution

import "testing"

func TestMySqrt(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
	}

	for _, test := range tests {
		ans := mySqrt(test.input)

		if ans != test.want {
			t.Fatalf("%d: expect %d to equal %d", test.input, ans, test.want)
		}
	}
}
