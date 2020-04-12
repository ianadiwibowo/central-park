package number_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ianadiwibowo/central-park/number"
)

func TestConvDecimalToBinary(t *testing.T) {
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
		if cmp.Equal(result, binaries[i]) == false {
			t.Errorf("Expected from %v: %v. Got: %v", d, binaries[i], result)
		}
	}
}

func TestConvDecimalToNegaBinary(t *testing.T) {
	decimals := []int{
		85,
		52,
		-5,
		-4,
		-3,
		-2,
		-1,
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		25,
	}
	negaBinaries := [][]int{
		[]int{1, 0, 1, 0, 1, 0, 1},
		[]int{1, 1, 1, 0, 1, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 0, 0},
		[]int{1, 1, 0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0},
		[]int{1},
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{1, 0, 0},
		[]int{1, 0, 1},
		[]int{1, 1, 0, 1, 0},
		[]int{1, 1, 0, 1, 0, 0, 1},
	}

	for i, d := range decimals {
		result := number.ConvDecimalToNegaBinary(d)
		if cmp.Equal(result, negaBinaries[i]) == false {
			t.Errorf("Expected from %v: %v. Got: %v", d, negaBinaries[i], result)
		}
	}
}

func TestConvBinaryToDecimal(t *testing.T) {
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

func TestConvNegaBinaryToDecimal(t *testing.T) {
	negaBinaries := [][]int{
		[]int{1, 0, 1, 0, 1, 0, 1},
		[]int{1, 1, 1, 0, 1, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 0, 0},
		[]int{1, 1, 0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0},
		[]int{1},
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{1, 0, 0},
		[]int{1, 0, 1},
		[]int{1, 1, 0, 1, 0},
		[]int{1, 1, 0, 1, 0, 0, 1},
	}
	decimals := []int{
		85,
		52,
		-5,
		-4,
		-3,
		-2,
		-1,
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		25,
	}

	for i, b := range negaBinaries {
		result := number.ConvNegaBinaryToDecimal(b)
		if result != decimals[i] {
			t.Errorf("Expected from %v: %v. Got: %v", b, decimals[i], result)
		}
	}
}

func TestAddBinary(t *testing.T) {
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
		if cmp.Equal(result, expected[i]) == false {
			t.Errorf("Expected from %v + %v: %v. Got: %v", v, b[i], expected[i], result)
		}
	}
}

func TestPadLeftWithZero(t *testing.T) {
	a := []int{1, 2, 3, 4}
	zeroCount := 5
	expected := []int{0, 0, 0, 0, 0, 1, 2, 3, 4}
	result := number.PadLeftWithZero(a, zeroCount)

	if cmp.Equal(result, expected) == false {
		t.Errorf("Expected: %v. Got: %v", expected, result)
	}
}

func TestAddBinaryDigit(t *testing.T) {
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

func TestFactorial(t *testing.T) {
	a_cases := []int{0, 1, 2, 3, 4, 5, 6}
	expectedResults := []int{1, 1, 2, 6, 24, 120, 720}

	for i, v := range a_cases {
		result := number.Factorial(v)
		if result != expectedResults[i] {
			t.Errorf("Expected from %v: %v. Got: %v", v, expectedResults[i], result)
		}
	}
}

func TestBinomialCoefficient(t *testing.T) {
	n_cases := []int{5, 3, 15, 5, 10, 0, 7}
	k_cases := []int{2, 2, 4, 4, 3, 2, 7}
	expectedResults := []int{10, 3, 1365, 5, 120, 0, 1}

	for i, v := range n_cases {
		result := number.BinomialCoefficient(v, k_cases[i])
		if result != expectedResults[i] {
			t.Errorf("Expected from %v, %v: %v. Got: %v", v, k_cases[i], expectedResults[i], result)
		}
	}
}
