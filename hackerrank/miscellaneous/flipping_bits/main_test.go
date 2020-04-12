package main

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestFlippingBits(t *testing.T) {
	assert(t, 9, 4294967286)
	assert(t, 2147483647, 2147483648)
	assert(t, 1, 4294967294)
	assert(t, 0, 4294967295)
	assert(t, 4, 4294967291)
	assert(t, 123456, 4294843839)
	assert(t, 802743475, 3492223820)
	assert(t, 35601423, 4259365872)
	assert(t, 4294967295, 0)
}

func assert(t *testing.T, input int64, expected int64) {
	testifyAssert.Equal(t, expected, FlippingBits(input))
}
