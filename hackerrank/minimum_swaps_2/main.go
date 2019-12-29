package main

import "fmt"

// https://www.hackerrank.com/challenges/minimum-swaps-2/problem
// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swapCount int32
	for n := 0; n < len(arr)-1; n++ {
		for arr[n] != int32(n+1) {
			temp := arr[n]
			arr[n] = arr[arr[n]-1]
			arr[temp-1] = temp
			swapCount++
		}
	}

	return swapCount
}

func main() {
	arr_cases := [][]int32{
		{7, 1, 3, 2, 4, 5, 6},
		{4, 3, 1, 2},
		{2, 3, 4, 1, 5},
		{1, 3, 5, 2, 4, 6, 7},
		{7, 6, 5, 4, 3, 2, 1},
	}
	expectations := []int32{
		5,
		3,
		3,
		3,
		3,
	}

	for i, v := range arr_cases {
		result := minimumSwaps(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
