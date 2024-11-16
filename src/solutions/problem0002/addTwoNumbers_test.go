package problem0002

import (
	"testing"

	"leetcode-go/src/structures"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums1, nums2, expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := [3]Case{
		{
			nums1:    []int{2, 4, 3},
			nums2:    []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			nums1:    []int{0},
			nums2:    []int{0},
			expected: []int{0},
		},
		{
			nums1:    []int{9, 9, 9, 9, 9, 9, 9},
			nums2:    []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, c := range cases {
		t.Run("2.两数相加", func(t *testing.T) {
			assert.Equal(
				t,
				structures.Int2List(c.expected),
				addTwoNumbers(
					structures.Int2List(c.nums1),
					structures.Int2List(c.nums2),
				),
			)
		})
	}
}
