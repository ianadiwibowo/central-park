package main

import (
	"fmt"
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestPrimality_GivenNonPrimeNumbers_ReturnsNotPrime(t *testing.T) {
	nonPrimeNumbers := []int32{0, 1, 4, 6, 8, 9, 10, 4, 6, 8, 9, 10, 12, 14, 15,
		16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 35, 36, 38, 39,
		40, 42, 44, 45, 46, 48, 49, 50, 51, 52, 54, 55, 56, 57, 58, 60, 62, 63,
		64, 65, 66, 68, 69, 70, 72, 74, 75, 76, 77, 78, 80, 81, 82, 84, 85, 86,
		87, 88, 90, 91, 92, 93, 94, 95, 96, 98, 99, 100, 121, 143, 144, 169, 225,
		289, 7917, 62678889, 982451651}

	for _, n := range nonPrimeNumbers {
		assert(t, n, "Not prime")
	}
}

func TestPrimality_GivenPrimeNumbers_ReturnsPrime(t *testing.T) {
	primeNumbers := []int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127,
		131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

	for _, n := range primeNumbers {
		assert(t, n, "Prime")
	}
}

func assert(t *testing.T, number int32, expected string) {
	testifyAssert.Equal(
		t,
		expected,
		Primality(number),
		fmt.Sprintf("Input number = %v\n", number),
	)
}
