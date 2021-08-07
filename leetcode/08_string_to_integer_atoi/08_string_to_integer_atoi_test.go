package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	a.Equal(t, 0, myAtoi(""))
	a.Equal(t, 0, myAtoi(" "))
	a.Equal(t, 0, myAtoi("  "))

	a.Equal(t, 0, myAtoi("   -"))
	a.Equal(t, 0, myAtoi("   - abc"))
	a.Equal(t, 0, myAtoi("   --"))
	a.Equal(t, 0, myAtoi("   -+12"))
	a.Equal(t, 0, myAtoi("        +"))
	a.Equal(t, 0, myAtoi("        +abc"))

	a.Equal(t, 42, myAtoi("42"))
	a.Equal(t, 42, myAtoi("+42"))
	a.Equal(t, -42, myAtoi("   -42"))
	a.Equal(t, 42, myAtoi("   +42"))
	a.Equal(t, 4193, myAtoi("4193 with words"))

	a.Equal(t, 2147483647, myAtoi("2147483648"))
	a.Equal(t, 2147483647, myAtoi("9999999999"))
	a.Equal(t, -2147483648, myAtoi("-2147483649"))
	a.Equal(t, -2147483648, myAtoi("-9999999999"))
	a.Equal(t, 2147483647, myAtoi(" 2147483648lol"))
	a.Equal(t, -2147483648, myAtoi(" -2147483649lol"))

	a.Equal(t, 2147483647, myAtoi("20000000000000000000"))
	a.Equal(t, -2147483648, myAtoi("-20000000000000000000"))

	a.Equal(t, 12345678, myAtoi("  0000000000012345678"))
	a.Equal(t, 12345678, myAtoi("  +0000000000012345678"))
	a.Equal(t, -12345678, myAtoi("  -0000000000012345678"))
}
