package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

// https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	anagramCount := make(map[string]int)

	for l := 1; l < len(s); l++ {
		for i := 0; i < len(s); i++ {
			if i+l-1 < len(s) {
				sub := s[i : i+l]
				sortedSub := strings.Split(sub, "")
				sort.Strings(sortedSub)
				anagram := strings.Join(sortedSub, "")
				anagramCount[anagram]++
			}
		}
	}

	var pairCount int32
	for _, v := range anagramCount {
		if v > 1 {
			x := big.NewInt(0)
			x.Binomial(int64(v), int64(2))
			pairCount += int32(x.Int64())
		}
	}

	return pairCount
}

func main() {
	s_cases := []string{
		"cdcd",
	}
	expectations := []int32{
		5,
	}

	for i, v := range s_cases {
		result := sherlockAndAnagrams(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
