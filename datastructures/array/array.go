package array

import (
	"sort"
)

// Equal checks whether two slices a and b are identical (true) or not (false)
// If importing external library is allowed, use cmp.Equal instead
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Mean calculates the average of array a
func Mean(a []int) float64 {
	var sum float64
	for _, v := range a {
		sum += float64(v)
	}
	return sum / float64(len(a))
}

// Median calculates the middle value of array a
func Median(a []int) float64 {
	sort.Ints(a)

	if len(a)%2 == 1 {
		midIndex := len(a) / 2
		return float64(a[midIndex])
	} else {
		midIndex1 := len(a)/2 - 1
		midIndex2 := midIndex1 + 1
		sum := a[midIndex1] + a[midIndex2]
		return float64(sum) / 2
	}
}

// Min finds the minimum value of array a
func Min(a []int) (minValue int) {
	if len(a) == 0 {
		return 0
	}

	minValue = a[0]
	for _, v := range a {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}

// Max finds the maximum value of array a
func Max(a []int) (maxValue int) {
	if len(a) == 0 {
		return 0
	}

	maxValue = a[0]
	for _, v := range a {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

// Reverse returns the array a in reversed order
func Reverse(a []int) []int {
	reversed := make([]int, len(a))
	copy(reversed, a)

	for i := 0; i < len(reversed)/2; i++ {
		temp := reversed[i]
		reversed[i] = reversed[len(a)-1-i]
		reversed[len(a)-1-i] = temp
	}

	return reversed
}

// Uniq returns distinct elements from the array a while preserving order
func Uniq(a []int) (result []int) {
	if len(a) == 0 {
		return []int{}
	}

	distinct := make(map[int]bool)
	for _, v := range a {
		distinct[v] = true
	}

	// This is needed to preserve order
	for _, v := range a {
		if distinct[v] {
			result = append(result, v)
			delete(distinct, v)
		}
	}
	return result
}

// Union merges set a and b while preserving order
func Union(a, b []int) (result []int) {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}

	counter := make(map[int]bool)
	for _, v := range a {
		counter[v] = true
	}
	for _, v := range b {
		counter[v] = true
	}

	// This is needed to preserve order
	for _, v := range a {
		if counter[v] {
			result = append(result, v)
			delete(counter, v)
		}
	}
	for _, v := range b {
		if counter[v] {
			result = append(result, v)
			delete(counter, v)
		}
	}
	return result
}

// Intersection finds distinct common elements between a and b while preserving order
func Intersection(a, b []int) (result []int) {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}

	common := make(map[int]bool)
	result = []int{}

	for _, v := range b {
		common[v] = true
	}

	for _, v := range a {
		if common[v] {
			result = append(result, v)
			delete(common, v)
		}
	}
	return result
}

// Difference finds distinct uncommon elements between a and b while preserving order
func Difference(a, b []int) (result []int) {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}

	common := make(map[int]int)
	result = []int{}
	const (
		NewElement    = 0
		CommonElement = 1
		OnlyInA       = 2
		OnlyInB       = 3
	)

	for _, v := range a {
		common[v] = OnlyInA
	}
	for _, v := range b {
		if common[v] == OnlyInA {
			common[v] = CommonElement
		} else if common[v] == NewElement {
			common[v] = OnlyInB
		}
	}

	// This is needed to preserve order
	for _, v := range a {
		if common[v] == OnlyInA {
			result = append(result, v)
			delete(common, v)
		}
	}
	for _, v := range b {
		if common[v] == OnlyInB {
			result = append(result, v)
			delete(common, v)
		}
	}

	return result
}

// TODO
func Permutation(a []int) (result [][]int) {
	return [][]int{}
}

// TODO
func Combination(a []int) (result [][]int) {
	return [][]int{}
}
