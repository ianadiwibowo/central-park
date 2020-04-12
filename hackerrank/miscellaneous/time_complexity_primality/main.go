package main

import (
	"math"
)

// Primality checks whether number is prime using trial division algorithm
func Primality(number int32) string {
	if number < 2 || hasSmallFactor(number) || hasBigFactor(number) {
		return "Not prime"
	}
	return "Prime"
}

func hasSmallFactor(number int32) bool {
	smallPrimes := []int32{2, 3, 5, 7}
	for _, n := range smallPrimes {
		if number > n && number%n == 0 {
			return true
		}
	}
	return false
}

func hasBigFactor(number int32) bool {
	limit := int32(math.Sqrt(float64(number)))
	var i int32
	for i = 11; i <= limit; i += 2 {
		if number%i == 0 {
			return true
		}
	}
	return false
}
