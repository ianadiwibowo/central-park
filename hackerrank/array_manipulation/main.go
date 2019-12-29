package main

import "fmt"

// https://www.hackerrank.com/challenges/crush/problem
// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n)
	for _, v := range queries {
		start := v[0] - 1
		end := v[1]
		value := int64(v[2])

		arr[start] += value
		if end < n {
			arr[end] -= value
		}
	}

	var localSum, max int64
	for _, v := range arr {
		localSum += v
		if localSum > max {
			max = localSum
		}
	}

	return max
}

func main() {
	n_cases := []int32{
		10,
		5,
		10,
		3,
	}
	queries_cases := [][][]int32{
		{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1},
		},
		{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		},
		{
			{2, 6, 8},
			{3, 5, 7},
			{1, 8, 1},
			{5, 9, 15},
		},
		{
			{1, 1, 1000},
			{3, 3, 50},
			{2, 2, 49},
		},
	}
	expectations := []int64{
		10,
		200,
		31,
		1000,
	}

	for i, v := range n_cases {
		result := arrayManipulation(v, queries_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			queries_cases[i],
			expectations[i],
			result,
		)
	}
}
