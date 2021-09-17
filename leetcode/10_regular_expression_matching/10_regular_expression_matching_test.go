package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	a.True(t, isMatch("a", "a"))
	a.True(t, isMatch("d", "d"))

	a.True(t, isMatch("a", "."))
	a.True(t, isMatch("c", "."))

	a.False(t, isMatch("aa", "p"))
	a.True(t, isMatch("ella", "ella"))
	a.True(t, isMatch("umbrella", "umbrella"))
	a.False(t, isMatch("la", "lax"))
	a.False(t, isMatch("lax", "la"))

	a.True(t, isMatch("aa", "a*"))
	a.True(t, isMatch("aabaa", "a*ba*"))

	a.True(t, isMatch("ab", ".*"))
}
