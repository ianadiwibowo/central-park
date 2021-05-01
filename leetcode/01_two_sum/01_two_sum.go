package main

import "fmt"

// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("Result: %v, Expected: %v\n", twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	fmt.Printf("Result: %v, Expected: %v\n", twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	fmt.Printf("Result: %v, Expected: %v\n", twoSum([]int{3, 3}, 6), []int{0, 1})
	fmt.Printf("Result: %v, Expected: %v\n", twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19), []int{8, 9})
	fmt.Printf("Result: %v, Expected: %v\n", twoSum([]int{2, 5, 5, 10}, 10), []int{1, 2})
}
