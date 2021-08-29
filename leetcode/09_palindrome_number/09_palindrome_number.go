package main

// 9. Palindrome Number (Easy)
// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// splitNumberIntoDigits
	digits := []int{}
	count := 0
	for x != 0 {
		currentDigit := x % 10
		digits = append(digits, currentDigit)
		x /= 10
		count += 1
	}

	// checkForPalindrome()
	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}
