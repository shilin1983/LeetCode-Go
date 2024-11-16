package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	cases := [3]Case{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, c := range cases {
		t.Run("1.两数之和", func(t *testing.T) {
			assert.Equal(
				t,
				c.expected,
				twoSum(c.nums, c.target),
			)
		})
	}
}
