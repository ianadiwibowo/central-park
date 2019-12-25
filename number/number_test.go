package number_test

import (
	"testing"

	number "github.com/ianadiwibowo/central-park/number"
)

func TestNumberConvDecimalToBinary(t *testing.T) {
	decimals := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		17,
		29,
		256,
		25000,
		1000000000,
	}
	binaries := [][]int{
		[]int{0},
		[]int{1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, 0, 0},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{1, 0, 0, 0},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 1, 1, 0, 1},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0},
		[]int{1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for i, d := range decimals {
		result := number.ConvDecimalToBinary(d)
		if equal(result, binaries[i]) == false {
			t.Errorf("Expected from %v: %v. Got: %v", d, binaries[i], result)
		}
	}
}

func TestNumberConvBinaryToDecimal(t *testing.T) {
	binaries := [][]int{
		[]int{0},
		[]int{1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, 0, 0},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{1, 0, 0, 0},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 1, 1, 0, 1},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0},
		[]int{1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	decimals := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		17,
		29,
		256,
		25000,
		1000000000,
	}

	for i, b := range binaries {
		result := number.ConvBinaryToDecimal(b)
		if result != decimals[i] {
			t.Errorf("Expected from %v: %v. Got: %v", b, decimals[i], result)
		}
	}
}

func TestNumberAddBinary(t *testing.T) {
	a := [][]int{
		[]int{0},
		[]int{0},
		[]int{1},
		[]int{1},
		[]int{1, 1, 0, 1},
		[]int{1, 0, 1, 0, 1, 0, 1, 0},
		[]int{1, 0, 1, 0},
		[]int{1},
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1},
	}
	b := [][]int{
		[]int{0},
		[]int{1},
		[]int{0},
		[]int{1},
		[]int{1, 0, 1, 1, 1},
		[]int{1, 1, 0, 0, 1, 1, 0, 0},
		[]int{1, 1, 0, 0, 1, 1, 0, 0},
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		[]int{1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1},
	}
	expected := [][]int{
		[]int{0},
		[]int{1},
		[]int{1},
		[]int{1, 0},
		[]int{1, 0, 0, 1, 0, 0},
		[]int{1, 0, 1, 1, 1, 0, 1, 1, 0},
		[]int{1, 1, 0, 1, 0, 1, 1, 0},
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		[]int{1, 0, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 0},
	}

	for i, v := range a {
		result := number.AddBinary(v, b[i])
		if equal(result, expected[i]) == false {
			t.Errorf("Expected from %v + %v: %v. Got: %v", v, b[i], expected[i], result)
		}
	}
}

func TestNumberPadLeftWithZero(t *testing.T) {
	a := []int{1, 2, 3, 4}
	zeroCount := 5
	expected := []int{0, 0, 0, 0, 0, 1, 2, 3, 4}
	result := number.PadLeftWithZero(a, zeroCount)

	if equal(result, expected) == false {
		t.Errorf("Expected: %v. Got: %v", expected, result)
	}
}

func TestNumberAddBinaryDigit(t *testing.T) {
	a := []int{0, 0, 0, 0, 1, 1, 1, 1}
	b := []int{0, 0, 1, 1, 0, 0, 1, 1}
	carries := []int{0, 1, 0, 1, 0, 1, 0, 1}
	expectedResults := []int{0, 1, 1, 0, 1, 0, 0, 1}
	expectedCarryOuts := []int{0, 0, 0, 1, 0, 1, 1, 1}

	for i, v := range a {
		result, carryOut := number.AddBinaryDigit(v, b[i], carries[i])
		if result != expectedResults[i] && carryOut != expectedCarryOuts[i] {
			t.Errorf("Expected from %v + %v: %v and %v. Got: %v and %v",
				v,
				b[i],
				expectedResults[i],
				expectedCarryOuts[i],
				result,
				carryOut,
			)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
