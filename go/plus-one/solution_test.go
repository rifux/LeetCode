package solution

import "testing"

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, test := range tests {
		ans := plusOne(test.input)
		if len(ans) == len(test.want) {
			for index, _ := range test.want {
				if ans[index] != test.want[index] {
					t.Fatalf("%v: expected %v to equal %v", test.input, ans, test.want)
				}
			}
		} else {
			t.Fatalf("%v: answer length %d expected to equal %d\n%v != %v", test.input, len(ans), len(test.want), ans, test.want)
		}
	}
}
