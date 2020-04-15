package main

import "fmt"

// https://www.hackerrank.com/challenges/sock-merchant/problem
// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	colorCount := make(map[int32]int32)
	for _, v := range ar {
		colorCount[v] += 1
	}

	var pairs int32
	for _, value := range colorCount {
		pairs += value / 2
	}

	return pairs
}

func main() {
	n_cases := []int32{
		9,
		10,
		1,
		20,
		21,
		9,
		8,
	}
	ar_cases := [][]int32{
		[]int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
		[]int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3},
		[]int32{500},
		[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 20},
		[]int32{1, 1, 1, 2, 2, 2, 3, 3, 3},
		[]int32{1, 1, 1, 1, 2, 2, 2, 2},
	}
	expectations := []int32{
		3,
		4,
		0,
		0,
		1,
		3,
		4,
	}

	for i, v := range n_cases {
		result := sockMerchant(v, ar_cases[i])
		fmt.Printf("Passed: %v. Test: %v, %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			ar_cases[i],
			expectations[i],
			result,
		)
	}
}
