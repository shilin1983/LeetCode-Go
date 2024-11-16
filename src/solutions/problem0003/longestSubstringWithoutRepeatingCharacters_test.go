package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case = struct {
	s        string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := [3]Case{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
	}

	for _, c := range cases {
		t.Run("3.无重复字符的最长子串", func(t *testing.T) {
			assert.Equal(
				t,
				c.expected,
				lengthOfLongestSubstring(c.s),
			)
		})
	}
}
