package main

import "fmt"

// https://www.hackerrank.com/challenges/ctci-bubble-sort/problem
// Complete the countSwaps function below.
func countSwaps(a []int32) string {
	swapCount := 0
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				swapCount++
			}
		}
	}

	return fmt.Sprintf("Array is sorted in %v swaps. First Element: %v Last Element: %v", swapCount, a[0], a[len(a)-1])
}

func main() {
	a_cases := [][]int32{
		[]int32{6, 4, 1},
		[]int32{1, 2, 3},
		[]int32{3, 2, 1},
	}
	expectations := []string{
		"Array is sorted in 3 swaps. First Element: 1 Last Element: 6",
		"Array is sorted in 0 swaps. First Element: 1 Last Element: 3",
		"Array is sorted in 3 swaps. First Element: 1 Last Element: 3",
	}

	for i, v := range a_cases {
		result := countSwaps(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
