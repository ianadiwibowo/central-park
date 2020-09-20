package sort_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 5, 2, 8, 9, 3, 3, 0, 100}
	expected := []int{0, 1, 2, 3, 3, 5, 8, 9, 100}

	assert.Equal(t, expected, sort.QuickSort(a))
}

func TestQuickSort_WithEmptyArray(t *testing.T) {
	a := []int{}
	expected := []int{}

	assert.Equal(t, expected, sort.QuickSort(a))
}

func TestQuickSort_WithSingleValue(t *testing.T) {
	a := []int{12}
	expected := []int{12}

	assert.Equal(t, expected, sort.QuickSort(a))
}

func TestQuickSort_WithWorstCaseScenario(t *testing.T) {
	a := []int{999, 888, 777, 666, 555, 444, 333, 222}
	expected := []int{222, 333, 444, 555, 666, 777, 888, 999}

	assert.Equal(t, expected, sort.QuickSort(a))
}

func TestQuickSort_WithNegativeValue(t *testing.T) {
	a := []int{5, 3, -1, 2, -8, -9, 0}
	expected := []int{-9, -8, -1, 0, 2, 3, 5}

	assert.Equal(t, expected, sort.QuickSort(a))
}

func TestQuickSort_WithAllSameValues(t *testing.T) {
	a := []int{8, 8, 8, 8}
	expected := []int{8, 8, 8, 8}

	assert.Equal(t, expected, sort.QuickSort(a))
}
