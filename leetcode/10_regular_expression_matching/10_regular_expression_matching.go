package main

import "fmt"

// 10. Regular Expression Matching (Hard)
// https://leetcode.com/problems/regular-expression-matching/

func isMatch(s string, p string) bool {
	stringIndex := 0
	currentChar := ""

	patternIndex := 0
	charToMatch := ""
	starMode := false
	dotMode := false

	for stringIndex < len(s) || patternIndex < len(p) {
		fmt.Println("\npatternIndex:", patternIndex, "| stringIndex:", stringIndex) // DEBUG

		// Current pattern char
		if patternIndex < len(p) { // If pattern is still available
			if !starMode {
				charToMatch = string(p[patternIndex])
				if patternIndex+1 < len(p) && string(p[patternIndex+1]) == "*" {
					starMode = true
				}
			}
			fmt.Print("charToMatch: ", charToMatch, " | starMode: ", starMode) // DEBUG

			if charToMatch == "." {
				dotMode = true
			}
			fmt.Println(" | dotMode:", dotMode) // DEBUG
		} else { // If pattern is done but string is still available
			if stringIndex < len(s) {
				return false
			}
		}

		// Current string char
		if stringIndex < len(s) { // If string is still available
			currentChar = string(s[stringIndex])
		} else { // If string is done but pattern is still available
			currentChar = ""
		}
		fmt.Println("currentChar:", currentChar) // DEBUG

		// Compare pattern vs. string char
		if dotMode || currentChar == charToMatch { // Match
			fmt.Println("Match") // DEBUG
			stringIndex += 1     // Advance the string char
			if !starMode {
				patternIndex += 1 // Advance the pattern char by 1 if not star mode
			}
		} else { // Not match
			fmt.Println("Not Match") // DEBUG
			if starMode {
				// if patternIndex+2 >= len(p) {
				// 	return false
				// }
				patternIndex += 2 // Advance the pattern char by 2 if star mode
				starMode = false
			} else {
				if patternIndex+1 >= len(p) {
					return false
				}
				patternIndex += 1 // Advance the pattern char by 1 if not star mode
			}
			dotMode = false
		}
	}

	return true
}
