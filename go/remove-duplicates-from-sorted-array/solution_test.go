package solution

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, TT := range tests {
		ans := removeDuplicates(TT.input)
		if ans != TT.want {
			t.Fatalf("%v: want %v, not %v", TT.input, TT.want, ans)
		}
	}
}
