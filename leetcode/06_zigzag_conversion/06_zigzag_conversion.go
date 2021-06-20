package main

import (
	"strings"
)

// 6. Zigzag Conversion (Medium)
// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	arr := make([]string, numRows)
	currentChar := 0
	currentRow := 0
	currentCol := 0
	dotCaseAllowedRow := numRows - 2

	for currentChar < len(s) {
		if currentCol%(numRows-1) == 0 { // FULL Case
			arr[currentRow] = arr[currentRow] + string(s[currentChar])
			currentChar += 1
			dotCaseAllowedRow = numRows - 2
		} else { // DOT Case
			if currentRow == dotCaseAllowedRow {
				arr[currentRow] = arr[currentRow] + string(s[currentChar])
				currentChar += 1
			} else {
				arr[currentRow] = arr[currentRow] + " "
			}

			if currentRow+1 == numRows {
				dotCaseAllowedRow -= 1
			}
		}

		currentRow += 1
		if currentRow == numRows {
			currentRow = 0
			currentCol += 1
		}
	}

	result := ""
	for _, a := range arr {
		result += strings.Replace(a, " ", "", -1)
	}

	return result
}
