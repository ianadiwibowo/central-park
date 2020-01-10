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

func TestMean(t *testing.T) {
	a_cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{10, 5, 3, 2, 6, 8, 3, 7, 7, 1, 90},
	}
	expectations := []float64{
		4.0,
		12.909090909090908,
	}

	for i, v := range a_cases {
		result := array.Mean(v)
		if result != expectations[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectations[i], result)
		}
	}
}

func TestMedianWithOddAray(t *testing.T) {
	a_cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{10, 5, 3, 2, 6, 8, 3, 7, 7, 1, 90},
		{10, 20, 30},
	}
	expectations := []float64{
		4,
		6,
		20,
	}

	for i, v := range a_cases {
		result := array.Median(v)
		if result != expectations[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectations[i], result)
		}
	}
}

func TestMedianWithEvenAray(t *testing.T) {
	a_cases := [][]int{
		{1, 2, 4, 6},
		{4, 5, 6, 8, 9, 1, 2, 3},
		{3, 13, 7, 5, 21, 23, 23, 40, 23, 14, 12, 56, 23, 29},
	}
	expectations := []float64{
		3,
		4.5,
		22,
	}

	for i, v := range a_cases {
		result := array.Median(v)
		if result != expectations[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectations[i], result)
		}
	}
}

func TestMin(t *testing.T) {
	a_cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{10, 5, 3, 2, 6, -800, 3, 7, 7, 1, 90},
		{6, 6, 6},
		{-1},
		{},
	}
	expectations := []int{
		1,
		-800,
		6,
		-1,
		0,
	}

	for i, v := range a_cases {
		result := array.Min(v)
		if result != expectations[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectations[i], result)
		}
	}
}

func TestMax(t *testing.T) {
	a_cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{10, 5, 3, 2, 6, -800, 3, 7, 7, 1, 90},
		{6, 6, 6},
		{-1},
		{},
	}
	expectations := []int{
		7,
		90,
		6,
		-1,
		0,
	}

	for i, v := range a_cases {
		result := array.Max(v)
		if result != expectations[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectations[i], result)
		}
	}
}
