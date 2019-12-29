package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem
// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	remainder := d % int32(len(a))
	if remainder == 0 {
		return a
	}

	return append(a[remainder:], a[:remainder]...)
}

func main() {
	a_cases := [][]int32{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
		{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97},
	}
	d_cases := []int32{
		4,
		5,
		6,
		10,
		13,
	}
	expectations := [][]int32{
		{5, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 1},
		{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
		{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60},
	}

	for i, v := range a_cases {
		result := rotLeft(v, d_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			cmp.Equal(result, expectations[i]),
			v,
			d_cases[i],
			expectations[i],
			result,
		)
	}
}
