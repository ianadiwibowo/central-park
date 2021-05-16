package main

import (
	"strings"
)

// 3. Longest Substring Without Repeating Characters (Medium)
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	longestSubstringLength := 0
	currentSubstring := ""
	hashmap := make(map[string]int)

	for i, char := range s {
		char := string(char)
		_, charExists := hashmap[char]

		if !charExists { // New unique char in the substring
			hashmap[char] = i

			currentSubstring = currentSubstring + char
			if len(currentSubstring) > longestSubstringLength {
				longestSubstringLength = len(currentSubstring)
			}
		} else { // No longer unique
			currentSubstring = currentSubstring[strings.Index(currentSubstring, char)+1:] + char

			hashmap = make(map[string]int)
			start := i - len(currentSubstring)
			for _, newChar := range currentSubstring {
				hashmap[string(newChar)] = start
				start += 1
			}
		}
	}

	return longestSubstringLength
}
