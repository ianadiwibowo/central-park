package main

import (
	"fmt"
	"sort"
)

// https://www.hackerrank.com/challenges/mark-and-toys/problem
// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	var totalPrice, toysCount int32

	for _, v := range prices {
		if totalPrice+v > k {
			break
		}

		totalPrice += v
		toysCount++
	}

	return toysCount
}

func main() {
	prices_cases := [][]int32{
		{1, 2, 3, 4},
		{1, 12, 5, 111, 200, 1000, 10},
		{3, 7, 2, 9, 4},
	}
	k_cases := []int32{
		7,
		50,
		15,
	}
	expectations := []int32{
		3,
		4,
		3,
	}

	for i, v := range prices_cases {
		result := maximumToys(v, k_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			prices_cases[i],
			expectations[i],
			result,
		)
	}
}
