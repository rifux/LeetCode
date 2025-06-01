package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"]", false},
	}

	for _, TT := range tests {
		ans := isValid(TT.input)
		if ans != TT.want {
			t.Fatalf("%s: want %t, got %t", TT.input, TT.want, ans)
		}
	}
}
