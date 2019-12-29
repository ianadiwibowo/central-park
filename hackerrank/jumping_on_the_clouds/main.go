package main

import "fmt"

// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem
// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var currentPos, jumps int
	for currentPos < len(c)-1 {
		if (currentPos+2 < len(c)) && (c[currentPos+2] == 0) {
			currentPos++
		}
		currentPos++
		jumps++
	}

	return int32(jumps)
}

func main() {
	c_cases := [][]int32{
		[]int32{0, 0, 1, 0, 0, 1, 0},
		[]int32{0, 0, 0, 0, 1, 0},
		[]int32{0, 0, 0, 0, 0},
		[]int32{0, 0, 0, 0, 0, 1, 0},
		[]int32{0, 0, 1, 0, 1, 0, 0},
		[]int32{0, 0},
		[]int32{0, 1, 0},
	}
	expectations := []int32{
		4,
		3,
		2,
		3,
		4,
		1,
		1,
	}

	for i, v := range c_cases {
		result := jumpingOnClouds(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
