package array_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/array"
	"github.com/stretchr/testify/assert"
)

func TestEqual_GivenEqualInputs_ReturnsTrue(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	assert.True(t, array.Equal(a, b))
}

func TestEqual_GivenDifferentLength_ReturnsFalse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6}

	assert.False(t, array.Equal(a, b))
}

func TestEqual_GivenSameLength_ButDifferentInputs_ReturnsFalse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 6, 9, 0, 9}

	assert.False(t, array.Equal(a, b))
}

func TestMean(t *testing.T) {
	assert.Equal(t, 4.0,
		array.Mean([]int{1, 2, 3, 4, 5, 6, 7}),
	)
	assert.Equal(t, 12.909090909090908,
		array.Mean([]int{10, 5, 3, 2, 6, 8, 3, 7, 7, 1, 90}),
	)
}

func TestMedian_GivenOddLengthArray(t *testing.T) {
	assert.Equal(t, 4.0, array.Median([]int{1, 2, 3, 4, 5, 6, 7}))
	assert.Equal(t, 20.0, array.Median([]int{10, 20, 30}))
	assert.Equal(t, 6.0,
		array.Median([]int{10, 5, 3, 2, 6, 8, 3, 7, 7, 1, 90}),
	)
}

func TestMedian_GivenEvenLengthAray(t *testing.T) {
	assert.Equal(t, 3.0, array.Median([]int{1, 2, 4, 6}))
	assert.Equal(t, 4.5,
		array.Median([]int{4, 5, 6, 8, 9, 1, 2, 3}),
	)
	assert.Equal(t, 22.0,
		array.Median([]int{3, 13, 7, 5, 21, 23, 23, 40, 23, 14, 12,
			56, 23, 29}),
	)
}

func TestMin(t *testing.T) {
	assert.Equal(t, 0, array.Min([]int{}))
	assert.Equal(t, -1, array.Min([]int{-1}))
	assert.Equal(t, 6, array.Min([]int{6, 6, 6}))
	assert.Equal(t, 1, array.Min([]int{1, 2, 3, 4, 5, 6, 7}))
	assert.Equal(t, -800,
		array.Min([]int{10, 5, 3, 2, 6, -800, 3, 7, 7, 1, 90}),
	)
}

func TestMax(t *testing.T) {
	assert.Equal(t, 0, array.Max([]int{}))
	assert.Equal(t, -1, array.Max([]int{-1}))
	assert.Equal(t, 6, array.Max([]int{6, 6, 6}))
	assert.Equal(t, 7, array.Max([]int{1, 2, 3, 4, 5, 6, 7}))
	assert.Equal(t, 90,
		array.Max([]int{10, 5, 3, 2, 6, -800, 3, 7, 7, 1, 90}),
	)
}

func TestReverse(t *testing.T) {
	assert.Equal(t,
		[]int{},
		array.Reverse([]int{}),
	)
	assert.Equal(t,
		[]int{-1},
		array.Reverse([]int{-1}),
	)
	assert.Equal(t,
		[]int{6, 6, 6},
		array.Reverse([]int{6, 6, 6}),
	)
	assert.Equal(t,
		[]int{7, 6, 5, 4, 3, 2, 1},
		array.Reverse([]int{1, 2, 3, 4, 5, 6, 7}),
	)
	assert.Equal(t,
		[]int{10, 5, 3, 2, 6, -800, 3, 7, 7, 1, 90},
		array.Reverse([]int{90, 1, 7, 7, 3, -800, 6, 2, 3, 5, 10}),
	)
	assert.Equal(t,
		[]int{-4, -3, -2, -1},
		array.Reverse([]int{-1, -2, -3, -4}),
	)
	assert.Equal(t,
		[]int{-5, -4, -3, -2, -1},
		array.Reverse([]int{-1, -2, -3, -4, -5}),
	)
}

func TestUniq_GivenEmptyInput_ReturnsEmptyOutput(t *testing.T) {
	assert.Equal(t, []int{}, array.Uniq([]int{}))
}

func TestUniq_GivenDistinctInput_ReturnsInput(t *testing.T) {
	assert.Equal(t, []int{1}, array.Uniq([]int{1}))
	assert.Equal(t, []int{1, 2, 3}, array.Uniq([]int{1, 2, 3}))
}

func TestUniq_GivenNonDistinctInput_ReturnsDistinctOutput(t *testing.T) {
	assert.Equal(t, []int{1}, array.Uniq([]int{1, 1}))
	assert.Equal(t, []int{1, 2}, array.Uniq([]int{1, 1, 2, 2}))
	assert.Equal(t, []int{3, 1, 2}, array.Uniq([]int{3, 1, 1, 2, 2}))
	assert.Equal(t, []int{7, 6, 5, 10},
		array.Uniq([]int{7, 6, 7, 7, 5, 6, 10, 5, 10, 7}),
	)
	assert.Equal(t, []int{11, 10, 9, 1},
		array.Uniq([]int{11, 10, 11, 10, 9, 9, 1}),
	)
}

func TestUnion_GivenEmptyInputs_ReturnsEmptyOutput(t *testing.T) {
	assert.Equal(t, []int{},
		array.Union([]int{}, []int{}))
}

func TestUnion_GivenSameInputs(t *testing.T) {
	assert.Equal(t, []int{1},
		array.Union([]int{1}, []int{1}))
	assert.Equal(t, []int{1, 2},
		array.Union([]int{1, 2}, []int{1, 2}))
}

func TestUnion_GivenDifferentInputs(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4},
		array.Union([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3},
		array.Union([]int{1, 2}, []int{2, 3}))
	assert.Equal(t, []int{10, 8, 6, 4, 2, 9, 7, 5, 3, 1},
		array.Union([]int{10, 8, 6, 4, 2}, []int{9, 7, 5, 3, 1}))
	assert.Equal(t, []int{99, 4, 2, 3, 1},
		array.Union([]int{99, 4, 2}, []int{3, 99, 1, 4, 2}))
	assert.Equal(t, []int{1, 2, 3, 4, 5},
		array.Union([]int{1}, []int{1, 2, 3, 4, 5}))
}

func TestIntersection_GivenEmptyInputs_ReturnsEmptyOutput(t *testing.T) {
	assert.Equal(t, []int{},
		array.Intersection([]int{}, []int{}))
}

func TestIntersection_GivenSameInputs(t *testing.T) {
	assert.Equal(t, []int{1},
		array.Intersection([]int{1}, []int{1}))
	assert.Equal(t, []int{1, 2},
		array.Intersection([]int{1, 2}, []int{1, 2}))
}

func TestIntersection_GivenDifferentInputs(t *testing.T) {
	assert.Equal(t, []int{},
		array.Intersection([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{},
		array.Intersection([]int{1, 2, 3}, []int{4, 5, 6, 7, 8}))
}

func TestIntersection_GivenOverlappingInputs(t *testing.T) {
	assert.Equal(t, []int{1},
		array.Intersection([]int{1, 2}, []int{1, 3}))
	assert.Equal(t, []int{3, 4, 5},
		array.Intersection([]int{1, 2, 3, 4, 5}, []int{5, 4, 3}))
	assert.Equal(t, []int{1, 2, 3},
		array.Intersection([]int{8, 1, 2, 3}, []int{5, 4, 3, 2, 1, 9, 5, 4}))
	assert.Equal(t, []int{2},
		array.Intersection([]int{1, 1, 1, 2, 2, 2}, []int{2, 3, 3, 3, 3}))
}

func TestDifference_GivenEmptyInputs_ReturnsEmptyOutput(t *testing.T) {
	assert.Equal(t, []int{},
		array.Difference([]int{}, []int{}))
}

func TestDifference_GivenSameInputs_ReturnsEmptyOutput(t *testing.T) {
	assert.Equal(t, []int{},
		array.Difference([]int{1}, []int{1}))
	assert.Equal(t, []int{},
		array.Difference([]int{1, 2}, []int{1, 2}))
}

func TestDifference_GivenDifferentInputs(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4},
		array.Difference([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8},
		array.Difference([]int{1, 2, 2, 3}, []int{4, 5, 5, 6, 7, 8}))
}

func TestDifference_GivenOverlappingInputs(t *testing.T) {
	assert.Equal(t, []int{2, 3},
		array.Difference([]int{1, 1, 1, 2, 2, 2}, []int{1, 1, 1, 3, 3, 3}))
	assert.Equal(t, []int{1, 2},
		array.Difference([]int{1, 2, 3, 4, 5}, []int{5, 4, 3}))
	assert.Equal(t, []int{8, 5, 4, 9},
		array.Difference([]int{8, 1, 2, 3}, []int{5, 4, 3, 2, 1, 9, 5, 4}))
	assert.Equal(t, []int{1, 3},
		array.Difference([]int{1, 1, 1, 2, 2, 2}, []int{2, 3, 3, 3, 3}))
}
