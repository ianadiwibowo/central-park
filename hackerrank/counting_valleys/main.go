package main

import "fmt"

// https://www.hackerrank.com/challenges/counting-valleys/problem
// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var altitude, valleyCount int32
	var isInValley bool

	for _, v := range s {
		if string(v) == "D" {
			altitude--
		} else {
			altitude++
		}

		if altitude < 0 {
			isInValley = true
		} else if altitude > 0 {
			isInValley = false
		} else {
			if isInValley {
				valleyCount++
				isInValley = false
			}
		}
	}

	return valleyCount
}

func main() {
	n_cases := []int32{
		8,
		12,
		2,
		2,
		16,
		16,
		16,
	}
	s_cases := []string{
		"UDDDUDUU",
		"DDUUDDUDUUUD",
		"DU",
		"UD",
		"UDUDUDUDUDUDUDUD",
		"DUDUDUDUDUDUDUDU",
		"DUUDDUUDDUUDDUUD",
	}
	expectations := []int32{
		1,
		2,
		1,
		0,
		0,
		8,
		4,
	}

	for i, v := range n_cases {
		result := countingValleys(v, s_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			s_cases[i],
			expectations[i],
			result,
		)
	}
}
