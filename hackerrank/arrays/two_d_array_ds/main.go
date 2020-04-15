package main

import "fmt"

// https://www.hackerrank.com/challenges/2d-array/problem
// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var sums [4][4]int32
	maxSum := int32(-63) // The minimum sum is -9 for all 7 number = -63

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sums[i][j] = arr[i][j] +
				arr[i][j+1] +
				arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] +
				arr[i+2][j+1] +
				arr[i+2][j+2]
			if sums[i][j] > maxSum {
				maxSum = sums[i][j]
			}
		}
	}

	return maxSum
}

func main() {
	arr_cases := [][][]int32{
		{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
		{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		{
			{-9, -9, -9, -9, -9, -9},
			{-9, -9, -9, -9, -9, -9},
			{-9, -9, -9, -9, -9, -9},
			{-9, -9, -9, -9, -9, -9},
			{-9, -9, -9, -9, -9, -9},
			{-9, -9, -9, -9, -9, -9},
		},
		{
			{9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9, 9},
		},
	}
	expectations := []int32{
		7,
		28,
		19,
		-63,
		63,
	}

	for i, v := range arr_cases {
		result := hourglassSum(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
