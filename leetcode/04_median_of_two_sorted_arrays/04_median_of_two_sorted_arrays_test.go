package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	a.Equal(t, 2.00000, findMedianSortedArrays([]int{1, 3}, []int{2}))                          // Odd
	a.Equal(t, 2.00000, findMedianSortedArrays([]int{2}, []int{1, 3}))                          // Odd
	a.Equal(t, 2.50000, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))                       // Even
	a.Equal(t, 2.50000, findMedianSortedArrays([]int{3, 4}, []int{1, 2}))                       // Even
	a.Equal(t, 41.00000, findMedianSortedArrays([]int{10, 25, 55, 60}, []int{9, 27, 100, 200})) // Even
}
