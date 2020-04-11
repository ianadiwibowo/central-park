package main

import (
	"strings"
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestMakingAnagram_GivenNonAnagrams_ReturnsTotalAllChars(t *testing.T) {
	assert(t, "a", "b", 2)
	assert(t, "aaaaa", "bbbbb", 10)
}

func TestMakingAnagram_GivenValidInput_ReturnsCorrectly(t *testing.T) {
	assert(t, "a", "a", 0)
	assert(t, "a", "ab", 1)
	assert(t, "ab", "a", 1)
	assert(t, "abc", "a", 2)
	assert(t, "a", "abc", 2)
	assert(t, "aaa", "aab", 2)
	assert(t, "aaa", "baa", 2)
	assert(t, "abb", "acc", 4)
	assert(t, "abcdef", "abcdefg", 1)
	assert(t, "abcdef", "aghijklmnopqrstuvwxyz", 25)
	assert(t, "bcddddddddddddddddddddddddddddda", "a", 31)
	assert(t, "cde", "dcf", 2)
	assert(t, "cde", "abc", 4)
	assert(t, "abcdefghij", "klmnopqrstuvwxyzzzzzzzzzz", 35)
}

func TestMakingAnagram_GivenLargeInput_ReturnsCorrectly(t *testing.T) {
	a := strings.Repeat("a", 9999) + "c"
	b := strings.Repeat("b", 9999) + "c"
	assert(t, a, b, 9999+9999)
}

func assert(t *testing.T, a, b string, expected int32) {
	testifyAssert.Equal(t, expected, MakeAnagram(a, b))
}
