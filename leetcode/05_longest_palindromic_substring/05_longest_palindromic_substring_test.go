package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	a.Equal(t, "1", longestPalindrome("1"))
	a.Equal(t, "bab", longestPalindrome("babad"))
	a.Equal(t, "bb", longestPalindrome("cbbd"))
	a.Equal(t, "11", longestPalindrome("abcdefghijklmnopqrstuvwxyz11"))
	a.Equal(t, "11", longestPalindrome("11abcdefghijklmnopqrstuvwxyz"))
	a.Equal(t, "11", longestPalindrome("abcdefghijklmn11opqrstuvwxyz"))
}

func TestIsPalindrome(t *testing.T) {
	a.Equal(t, true, isPalindrome(""))
	a.Equal(t, true, isPalindrome("1"))
	a.Equal(t, true, isPalindrome("abcba"))
	a.Equal(t, true, isPalindrome("abba"))
	a.Equal(t, true, isPalindrome("1234567890987654321"))

	a.Equal(t, false, isPalindrome("ab"))
	a.Equal(t, false, isPalindrome("abc"))
	a.Equal(t, false, isPalindrome("abcdefghijklmnopqrstuvwxyz1234567890"))
}
