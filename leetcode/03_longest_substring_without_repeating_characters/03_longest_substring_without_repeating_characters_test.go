package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	a.Equal(t, 0, lengthOfLongestSubstring(""))

	a.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	a.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	a.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	a.Equal(t, 15, lengthOfLongestSubstring("abc123def456a789"))
	a.Equal(t, 12, lengthOfLongestSubstring("abc123def456f789"))

	a.Equal(t, 9, lengthOfLongestSubstring(" % a1234 %79a"))
}
