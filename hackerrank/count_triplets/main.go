package main

import (
	"fmt"
	"math/big"
)

// https://www.hackerrank.com/challenges/count-triplets-1/problem
// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	numCount := make(map[int64]int64)
	nextProgCount := make(map[int64]int64)
	var tripletCount int64
	for i := len(arr) - 1; i >= 0; i-- {
		numCount[arr[i]]++
		if r != 1 {
			nextProgCount[arr[i]] += numCount[arr[i]*r]
			tripletCount += nextProgCount[arr[i]*r]
		}
	}

	if r == 1 {
		for _, v := range numCount {
			x := big.NewInt(0)
			x.Binomial(int64(v), int64(3))
			tripletCount += x.Int64()
		}
	}

	return tripletCount
}

func main() {
	arr_cases := [][]int64{
		{1, 2, 2, 4},
		{1, 2, 4, 2},
		{1, 3, 9, 9, 27, 81},
		{1, 5, 5, 25, 125},
		{64, 1, 4, 16, 64},
		{1, 2, 1, 2, 4},
		{1, 1, 1},
		{1, 1, 1, 2, 2, 2, 3, 3, 3},
		{2, 18, 6, 1, 2, 3},
	}
	r_cases := []int64{
		2,
		2,
		3,
		5,
		4,
		2,
		1,
		1,
		3,
	}
	expectations := []int64{
		2,
		1,
		6,
		4,
		2,
		3,
		1,
		3,
		0,
	}

	for i, v := range arr_cases {
		result := countTriplets(v, r_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			r_cases[i],
			expectations[i],
			result,
		)
	}
}
