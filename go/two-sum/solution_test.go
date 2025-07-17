package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		name        string
		inputNums   []int
		inputTarget int
		want        []int
	}{
		{
			"Ascending",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			"Random",
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			"Equal",
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, twoSum(tc.inputNums, tc.inputTarget))
		})
	}
}
