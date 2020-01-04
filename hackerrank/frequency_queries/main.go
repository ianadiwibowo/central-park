package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// https://www.hackerrank.com/challenges/frequency-queries/problem
// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	data := make(map[int32]int32)
	counts := make(map[int32]int32)
	res := []int32{}

	for _, v := range queries {
		if v[0] == 1 { // Increment x
			if counts[data[v[1]]] > 0 {
				counts[data[v[1]]]--
			}
			data[v[1]]++
			counts[data[v[1]]]++
		} else if v[0] == 2 { // Decrement y
			if counts[data[v[1]]] > 0 {
				counts[data[v[1]]]--
			}
			if data[v[1]] > 0 {
				data[v[1]]--
			}
			counts[data[v[1]]]++
		} else { // Check count z
			if counts[v[1]] > 0 {
				res = append(res, 1)
			} else {
				res = append(res, 0)
			}
		}
	}

	return res
}

func main() {
	queries_cases := [][][]int32{
		{
			{1, 1},
			{2, 2},
			{3, 2},
			{1, 1},
			{1, 1},
			{2, 1},
			{3, 2},
		},
		{
			{1, 5},
			{1, 6},
			{3, 2},
			{1, 10},
			{1, 10},
			{1, 6},
			{2, 5},
			{3, 2},
		},
		{
			{3, 4},
			{2, 1003},
			{1, 16},
			{3, 1},
		},
		{
			{1, 3},
			{2, 3},
			{3, 2},
			{1, 4},
			{1, 5},
			{1, 5},
			{1, 4},
			{3, 2},
			{2, 4},
			{3, 2},
		},
	}
	expectations := [][]int32{
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1, 1},
	}

	for i, v := range queries_cases {
		result := freqQuery(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			cmp.Equal(result, expectations[i]),
			v,
			expectations[i],
			result,
		)
	}
}
