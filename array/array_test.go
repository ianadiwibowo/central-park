package array_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/array"
)

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
