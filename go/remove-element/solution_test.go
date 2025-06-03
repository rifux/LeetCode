package solution

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		inputSlice []int
		inputVal   int
		wantSlice  []int
		wantVal    int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 4, 0, 3},
			5,
		},
	}

	for _, TT := range tests {
		ans := removeElement(TT.inputSlice, TT.inputVal)

		newSlice := make([]int, ans)
		copy(newSlice, TT.inputSlice[:ans])

		if ans != TT.wantVal {
			t.Fatalf("%v: want %d, not %d", TT.inputSlice, TT.wantVal, ans)
		}

		sort.Ints(newSlice)
		sort.Ints(TT.wantSlice)
		for index, item := range newSlice {
			if item != TT.wantSlice[index] {
				t.Fatalf("at index %d: want %v, not %v", index, TT.wantSlice, newSlice)
			}
		}
	}
}
