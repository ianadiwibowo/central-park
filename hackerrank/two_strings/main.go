package main

import (
	"fmt"
	"strings"
)

// https://www.hackerrank.com/challenges/two-strings/problem
// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	for _, v := range s1 {
		if strings.Contains(s2, string(v)) {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	s1_cases := []string{
		"hello",
		"hi",
		"wouldyoulikefries",
		"hackerrankcommunity",
		"jackandjill",
		"writetoyourparents",
		"abcdefghijklm",
		"ab",
	}
	s2_cases := []string{
		"world",
		"world",
		"abcabcabcabcabcabc",
		"cdecdecdecde",
		"wentupthehill",
		"fghmqzldbc",
		"nopqrstuvwxyzabcd",
		"ccccccccccccba",
	}
	expectations := []string{
		"YES",
		"NO",
		"NO",
		"YES",
		"YES",
		"NO",
		"YES",
		"YES",
	}

	for i, v := range s1_cases {
		result := twoStrings(v, s2_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			s2_cases[i],
			expectations[i],
			result,
		)
	}
}
