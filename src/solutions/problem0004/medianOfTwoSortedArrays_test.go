package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums1, nums2 []int
	expected     float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	cases := [2]Case{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
	}

	for _, c := range cases {
		t.Run("4.寻找两个正序数组的中位数", func(t *testing.T) {
			assert.Equal(
				t,
				c.expected,
				findMedianSortedArrays(c.nums1, c.nums2),
			)
		})
	}
}
