package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	a.Equal(t, 42, myAtoi("42"))
	a.Equal(t, -42, myAtoi("   -42"))
	a.Equal(t, 4193, myAtoi("4193 with words"))
}
