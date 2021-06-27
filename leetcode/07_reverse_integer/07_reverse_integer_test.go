package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	a.Equal(t, 321, reverse(123))
	a.Equal(t, -321, reverse(-123))
	a.Equal(t, 21, reverse(120))
	a.Equal(t, 0, reverse(2147483647))
}

func TestReverseString(t *testing.T) {
	a.Equal(t, "a", reverseString("a"))
	a.Equal(t, "dcba", reverseString("abcd"))
	a.Equal(t, "edcba", reverseString("abcde"))
}
