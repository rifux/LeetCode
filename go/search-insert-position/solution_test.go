package solution

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		inputNums   []int
		inputTarget int
		want        int
	}{
		{[]int{1, 2, 5, 6}, 5, 2},
		{[]int{1, 2, 5, 6}, 2, 1},
		{[]int{1, 2, 5, 6}, 7, 4},
	}

	for _, TT := range tests {
		ans := searchInsert(TT.inputNums, TT.inputTarget)
		if ans != TT.want {
			t.Fatalf("nms: %v, trgt: %d:\nwant %d, not %d", TT.inputNums, TT.inputTarget, TT.want, ans)
		}
	}
}
