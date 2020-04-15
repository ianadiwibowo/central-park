package main

import (
	"fmt"
	"strings"
)

// https://www.hackerrank.com/challenges/repeated-string/problem
// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	slength := int64(len(s))

	quotient := n / slength
	remainder := n % slength

	aCountInS := int64(strings.Count(s, "a"))
	aCountInRemainingSubstring := int64(strings.Count(s[:remainder], "a"))

	return quotient*aCountInS + aCountInRemainingSubstring
}

func main() {
	s_cases := []string{
		"aba",
		"aba",
		"aba",
		"a",
		"abcac",
		"abcdefghij",
		"abcdefghij",
		"abcdefghij",
	}
	n_cases := []int64{
		10,
		11,
		12,
		1000000000000,
		10,
		9,
		2,
		11,
	}
	expectations := []int64{
		7,
		7,
		8,
		1000000000000,
		4,
		1,
		1,
		2,
	}

	for i, v := range s_cases {
		result := repeatedString(v, n_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			n_cases[i],
			expectations[i],
			result,
		)
	}
}
