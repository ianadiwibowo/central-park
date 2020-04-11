package main

import (
	"strings"
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestAlternatingCharacters_GivenEmptyString_ReturnsZero(t *testing.T) {
	assert(t, "", 0)
}

func TestAlternatingCharacters_GivenSingleChar_ReturnsZero(t *testing.T) {
	assert(t, "A", 0)
	assert(t, "B", 0)
}

func TestAlternatingCharacters_GivenTwoChars(t *testing.T) {
	assert(t, "AB", 0)
	assert(t, "BA", 0)
	assert(t, "AA", 1)
	assert(t, "BB", 1)
}

func TestAlternatingCharacters_GivenThreeChars(t *testing.T) {
	assert(t, "ABA", 0)
	assert(t, "BAB", 0)
	assert(t, "AAB", 1)
	assert(t, "ABB", 1)
	assert(t, "AAA", 2)
	assert(t, "BBB", 2)
}

func TestAlternatingCharacters_GivenLongerString(t *testing.T) {
	assert(t, "AAAA", 3)
	assert(t, "BBBBB", 4)
	assert(t, "ABABABAB", 0)
	assert(t, "BABABA", 0)
	assert(t, "AAABBB", 4)
}

func TestAlternatingCharacters_GivenVeryLongString(t *testing.T) {
	s := strings.Repeat("AB", 50000)
	assert(t, s, 0)

	s = strings.Repeat("A", 100000)
	assert(t, s, 99999)

	s = strings.Repeat("B", 100000)
	assert(t, s, 99999)

	s = strings.Repeat("A", 99999) + "B"
	assert(t, s, 99998)

	s = "A" + strings.Repeat("B", 99999)
	assert(t, s, 99998)
}

func assert(t *testing.T, s string, expected int32) {
	testifyAssert.Equal(t, expected, AlternatingCharacters(s))
}
