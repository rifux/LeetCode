package solution

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
	}

	for _, TT := range tests {
		ans := lengthOfLastWord(TT.input)
		if ans != TT.want {
			t.Fatalf("%s: want %d, got %d", TT.input, TT.want, ans)
		}
	}
}
