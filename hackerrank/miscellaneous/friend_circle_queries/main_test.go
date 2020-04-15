package main

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestMaxCircle_GivenNonOverlapFriendship(t *testing.T) {
	assert(t, [][]int32{{1, 2}}, []int32{2})
	assert(t, [][]int32{{1, 2}, {3, 4}}, []int32{2, 2})
	assert(t, [][]int32{{1, 2}, {3, 4}, {5, 6}}, []int32{2, 2, 2})
}

func TestMaxCircle_GivenOverlapFriendship(t *testing.T) {
	assert(t, [][]int32{{1, 2}, {1, 3}}, []int32{2, 3})
	assert(t, [][]int32{{1, 2}, {2, 3}, {3, 4}}, []int32{2, 3, 4})
	assert(t, [][]int32{{1, 2}, {3, 4}, {1, 3}}, []int32{2, 2, 4})
	assert(
		t,
		[][]int32{{1, 2}, {1, 3}, {4, 5}, {4, 6}, {2, 5}},
		[]int32{2, 3, 3, 3, 6},
	)
	assert(t,
		[][]int32{{1, 2}, {3, 4}, {2, 3}},
		[]int32{2, 2, 4},
	)
	assert(t,
		[][]int32{{1, 2}, {1, 3}},
		[]int32{2, 3},
	)
	assert(t,
		[][]int32{{1000000000, 23}, {11, 3778}, {7, 47}, {11, 1000000000}},
		[]int32{2, 2, 2, 4},
	)
	assert(t,
		[][]int32{{1, 2}, {3, 4}, {1, 3}, {5, 7}, {5, 6}, {7, 4}},
		[]int32{2, 2, 4, 4, 4, 7},
	)
}

func TestMaxCircle_GivenRepeatingQuery(t *testing.T) {
	assert(t, [][]int32{{1, 2}, {1, 2}}, []int32{2, 2})
	assert(
		t,
		[][]int32{{1, 2}, {3, 4}, {1, 3}, {1, 2}, {3, 4}, {1, 3}},
		[]int32{2, 2, 4, 4, 4, 4},
	)
}

func assert(t *testing.T, queries [][]int32, expected []int32) {
	testifyAssert.Equal(t, expected, MaxCircle(queries), queries)
}

// func TestMaxCircle_GivenVeryLargeQueries(t *testing.T) {
// 	assert(t,
// 		[][]int32{},
// 		[]int32{},
// 	)
// }
