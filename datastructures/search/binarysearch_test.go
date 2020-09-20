package search_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/search"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch_FindInTheExactMiddle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 5)
	assert.Equal(t, 5, index)
}

func TestBinarySearch_FindInTheLeftMost(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 1)
	assert.Equal(t, 1, index)
}

func TestBinarySearch_FindInTheRightMost(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 9)
	assert.Equal(t, 9, index)
}

func TestBinarySearch_FindInTheLeftArea(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 2)
	assert.Equal(t, 2, index)
}

func TestBinarySearch_FindInTheRightArea(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 7)
	assert.Equal(t, 7, index)
}

func TestBinarySearch_NotFound_WithValueInTheMiddle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 60, 70, 80, 90}
	index := search.BinarySearch(a, 6)
	assert.Equal(t, -1, index)
}

func TestBinarySearch_NotFound_WithValueLeanLeft(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, -5)
	assert.Equal(t, -1, index)
}

func TestBinarySearch_NotFound_WithValueLeanRight(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := search.BinarySearch(a, 15)
	assert.Equal(t, -1, index)
}
