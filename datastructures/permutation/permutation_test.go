package permutation_test

import (
	"fmt"
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/permutation"
	"github.com/stretchr/testify/assert"
)

func TestHeapPermutation_SizeOne(t *testing.T) {
	a := []int{5}
	expected := "[[5]]"

	assert.Equal(t, expected, fmt.Sprintf("%v", permutation.HeapPermutation(a)))
}

func TestHeapPermutation_SizeTwo(t *testing.T) {
	a := []int{1, 2}
	expected := "[[1 2] [2 1]]"

	assert.Equal(t, expected, fmt.Sprintf("%v", permutation.HeapPermutation(a)))
}

func TestHeapPermutation_SizeThree(t *testing.T) {
	a := []int{1, 2, 3}
	expected := "[[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]"

	assert.Equal(t, expected, fmt.Sprintf("%v", permutation.HeapPermutation(a)))
}

func TestHeapPermutation_SizeFour(t *testing.T) {
	a := []int{1, 2, 3, 4}
	expected := "[[1 2 3 4] [2 1 3 4] [3 1 2 4] [1 3 2 4] [2 3 1 4] [3 2 1 4] [4 2 3 1] [2 4 3 1] [3 4 2 1] [4 3 2 1] [2 3 4 1] [3 2 4 1] [4 1 3 2] [1 4 3 2] [3 4 1 2] [4 3 1 2] [1 3 4 2] [3 1 4 2] [4 1 2 3] [1 4 2 3] [2 4 1 3] [4 2 1 3] [1 2 4 3] [2 1 4 3]]"

	assert.Equal(t, expected, fmt.Sprintf("%v", permutation.HeapPermutation(a)))
}
