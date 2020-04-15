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
