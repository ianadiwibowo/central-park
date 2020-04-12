package main

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {
	assert(t, -7, -7)
}

func assert(t *testing.T, input int, expected int) {
	testifyAssert.Equal(t, expected, SomeFunction(input))
}
