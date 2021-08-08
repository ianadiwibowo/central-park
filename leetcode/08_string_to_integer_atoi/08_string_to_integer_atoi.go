package main

import (
	"fmt"
	"math"
)

// 8. String to Integer (atoi) (Medium)
// https://leetcode.com/problems/string-to-integer-atoi/

// Yeah this is a stupid puzzle
// I mean, why?
// I'm so upset that I'm not going to optimize these ugly codes
func myAtoi(s string) int {
	currentIndex := 0

	// Step 1: Leading space
	for _, c := range s {
		if string(c) == " " {
			currentIndex += 1
			continue
		} else {
			break
		}
	}

	// Step 2: Positive/negative
	if currentIndex == len(s) {
		return 0
	}

	isNegative := false
	if string(s[currentIndex]) == "-" {
		isNegative = true
		currentIndex += 1
	} else if string(s[currentIndex]) == "+" {
		currentIndex += 1
	}

	// Step 3: Leading zeros
	for i := currentIndex; i < len(s); i++ {
		if string(s[currentIndex]) == "0" {
			currentIndex += 1
			continue
		} else {
			break
		}
	}

	// Step 4: Real number conversion
	numberMap := make(map[string]int)
	for i := 0; i < 10; i++ {
		numberMap[fmt.Sprint(i)] = i
	}
	numberString := ""
	for i := currentIndex; i < len(s); i++ {
		_, ok := numberMap[string(s[i])]
		if !ok {
			break
		}
		numberString += string(s[i])
	}
	if len(numberString) == 0 {
		return 0
	}

	// Step 5: Min/max value
	MAX_VALUE := 2147483647
	MIN_VALUE := -2147483648
	if len(numberString) > 10 && isNegative {
		return MIN_VALUE
	}
	if len(numberString) > 10 && !isNegative {
		return MAX_VALUE
	}

	number := 0
	power := 0
	for i := len(numberString) - 1; i >= 0; i-- {
		number += numberMap[string(numberString[i])] * int(math.Pow(10, float64(power)))
		if isNegative && number > MAX_VALUE+1 {
			return MIN_VALUE
		}
		if !isNegative && number > MAX_VALUE {
			return MAX_VALUE
		}
		power += 1
	}
	if isNegative {
		number *= -1
	}

	return number
}
