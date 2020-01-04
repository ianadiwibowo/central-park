package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/common-child/problem
// Complete the commonChild function below.
var hash map[string]int

func commonChild(s1 string, s2 string) int32 {
	hash = make(map[string]int)

	return -777
}

func rCommonChild(s1 string, s2 string) {
	// Recurrence definition

	// Check cache first

	// Compute otherwise
	rCommonChild(s1 string, s2 string)
}

func main() {
	s1_cases := []string{
		"ABCD", // ABCD, ABC (012), ABD (013), ACD (023), BCD (123), AB, AC, AD, BC, BD, CD, A, B, C, D
		"HARRY",
		"AA",
		"SHINCHAN", // H N H A N
		"ABCDEF",
	}
	s2_cases := []string{
		"ABDC", // ABDC, ABD, ABC, ADC, BDC, AB, AD, AC, BD, BC, EC, A, B, C, D
		"SALLY",
		"BB",
		"NOHARAAA", // N H A A A A
		"FBDAMN",
	}
	expectations := []int32{
		3,
		2,
		0,
		3,
		2,
	}

	for i, v := range s1_cases {
		result := commonChild(v, s2_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			s2_cases[i],
			expectations[i],
			result,
		)
	}
}
