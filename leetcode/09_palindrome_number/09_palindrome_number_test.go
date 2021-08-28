package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a.Equal(t, false, isPalindrome(-121))
	a.Equal(t, false, isPalindrome(-2147483647))

	a.Equal(t, true, isPalindrome(1))
	a.Equal(t, true, isPalindrome(0))

	a.Equal(t, false, isPalindrome(10))
	a.Equal(t, true, isPalindrome(121))
	a.Equal(t, false, isPalindrome(12345))
	a.Equal(t, true, isPalindrome(543212345))
	a.Equal(t, false, isPalindrome(2147483647))
	a.Equal(t, true, isPalindrome(2147447412))
}
