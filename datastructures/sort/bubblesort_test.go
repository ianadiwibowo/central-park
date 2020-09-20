package sort_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ianadiwibowo/central-park/datastructures/sort"
)

func TestBubbleSort(t *testing.T) {
	a_cases := [][]int{
		{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{-9},
		{},
	}
	expectations := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		{-9},
		{},
	}
	for i, v := range a_cases {
		result := sort.BubbleSort(v)
		if cmp.Equal(result, expectations[i]) != true {
			t.Errorf("Test: %v. Expected: %v. Got: %v.", v, expectations[i], result)
		}
	}
}
