package main

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestCalculateCost(t *testing.T) {
	testifyAssert.Equal(t, int32(1),
		calculateCost(
			[][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]int32{{2, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		),
	)
	testifyAssert.Equal(t, int32(1),
		calculateCost(
			[][]int32{{2, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		),
	)
	testifyAssert.Equal(t, int32(9),
		calculateCost(
			[][]int32{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			[][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		),
	)
}

func TestFormingMagicSquare(t *testing.T) {
	assert(t, [][]int32{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 5},
	},
		1,
	)
}

func assert(t *testing.T, input [][]int32, expected int32) {
	testifyAssert.Equal(t, expected, FormingMagicSquare(input))
}

// func TestIsMagicSquare(t *testing.T) {
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}}))
// 	testifyAssert.True(t,
// 		isMagicSquare([][]int32{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}}))
// }
