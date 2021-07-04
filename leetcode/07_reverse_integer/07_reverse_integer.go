package main

import (
	"fmt"
	"strconv"
)

// 7. Reverse Integer (Medium)
// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	MAX_VALUE := 2147483647
	MIN_VALUE := -2147483648

	number := reverseString(fmt.Sprint(x))

	if string(number[len(number)-1]) == "-" {
		number = "-" + number[:len(number)-1]
	}

	result, err := strconv.Atoi(number)
	if err != nil || result > MAX_VALUE || result < MIN_VALUE {
		return 0
	}

	return result
}

func reverseString(s string) string {
	result := ""
	for _, c := range s {
		result = string(c) + result
	}
	return result
}
