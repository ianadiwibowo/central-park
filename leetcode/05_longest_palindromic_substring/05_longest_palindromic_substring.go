package main

// 5. Longest Palindromic Substring (Medium)
// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	for substringLength := len(s); substringLength > 1; substringLength-- {
		for startIndex := 0; startIndex <= len(s)-substringLength; startIndex++ {
			substring := s[startIndex : substringLength+startIndex]

			if isPalindrome(substring) {
				return substring
			}
		}
	}

	return string(s[0])
}

func isPalindrome(s string) bool {
	median := len(s) / 2

	for i := 0; i < median; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
