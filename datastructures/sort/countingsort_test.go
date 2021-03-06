package sort_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ianadiwibowo/central-park/datastructures/sort"
)

func TestCountingSort(t *testing.T) {
	a_cases := [][]int{
		{4, 0, 0, 1, 0, 2, 4, 5, 1},
		{0, 4, 1, 7, 5, 5, 6, 4, 3, 3, 4, 2, 1, 9, 8, 4, 6},
		{2, 1, 0},
		{5},
		{},
	}
	expectations := [][]int{
		{0, 0, 0, 1, 1, 2, 4, 4, 5},
		{0, 1, 1, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 7, 8, 9},
		{0, 1, 2},
		{5},
		{},
	}
	for i, v := range a_cases {
		result := sort.CountingSort(v)
		if cmp.Equal(result, expectations[i]) != true {
			t.Errorf("Test: %v. Expected: %v. Got: %v.", v, expectations[i], result)
		}
	}
}
