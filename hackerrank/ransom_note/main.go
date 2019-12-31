package main

import "fmt"

// https://www.hackerrank.com/challenges/ctci-ransom-note/problem
// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) string {
	wordCount := make(map[string]int)
	for _, v := range magazine {
		wordCount[v]++
	}

	for _, v := range note {
		if wordCount[v] == 0 {
			return "No"
		}
		wordCount[v]--
	}

	return "Yes"
}

func main() {
	magazine_cases := [][]string{
		{"give", "me", "one", "grand", "today", "night"},
		{"two", "times", "three", "is", "not", "four"},
		{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
		{"a", "a", "a", "b", "c", "d", "e", "e", "e"},
	}
	note_cases := [][]string{
		{"give", "one", "grand", "today"},
		{"two", "times", "two", "is", "four"},
		{"ive", "got", "some", "coconuts"},
		{"b", "b"},
	}
	expectations := []string{
		"Yes",
		"No",
		"No",
		"No",
	}

	for i, v := range magazine_cases {
		result := checkMagazine(v, note_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			note_cases[i],
			expectations[i],
			result,
		)
	}
}
