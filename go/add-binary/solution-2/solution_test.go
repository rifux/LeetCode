package solution2

import "testing"

/*
Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
*/

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		input [2]string
		want  string
	}{
		{[2]string{"11", "1"}, "100"},
		{[2]string{"1010", "1011"}, "10101"},
		{[2]string{"0", "1"}, "1"},
		{[2]string{"111", "1"}, "1000"},
		{[2]string{"1111", "1111"}, "11110"},
	}

	for _, test := range tests {
		ans := addBinary(test.input[0], test.input[1])
		if ans != test.want {
			t.Fatalf(`%v: expected "%s" to equal "%s"`, test.input, ans, test.want)
		}
	}
}
