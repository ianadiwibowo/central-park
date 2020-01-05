package array_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ianadiwibowo/central-park/array"
)

func TestBubbleSort(t *testing.T) {
	a := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := array.BubbleSort(a)
	if cmp.Equal(result, expected) != true {
		t.Errorf("Test: %v. Expected: %v. Got: %v.", a, expected, result)
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	if array.Equal(a, b) != true {
		t.Errorf("Expected: true. Got: false")
	}
}

func TestEqualWithDifferentLength(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6}

	if array.Equal(a, b) != false {
		t.Errorf("Expected: false. Got: true")
	}
}

func TestEqualWithSameLengthButNonIdenticalInputs(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 6, 9, 0, 9}

	if array.Equal(a, b) != false {
		t.Errorf("Expected: false. Got: true")
	}
}
