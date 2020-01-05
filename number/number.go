package number

import (
	"math"
	"sort"
)

// ConvDecimalToBinary converts decimal (base 10) to binary (base 2) in 0-1 array format
func ConvDecimalToBinary(value int) []int {
	if value == 0 {
		return []int{0}
	}

	binary := []int{}

	for value != 0 {
		binary = append([]int{value % 2}, binary...)
		value = value / 2
	}

	return binary
}

// ConvDecimalToNegaBinary converts decimal (base 10) to negative binary (base -2) in 0-1 array format
func ConvDecimalToNegaBinary(value int) []int {
	if value == 0 {
		return []int{0}
	}

	negaBinary := []int{}

	for value != 0 {
		remainder := value % -2
		value = value / -2
		if remainder < 0 {
			remainder += 2
			value += 1
		}
		negaBinary = append([]int{remainder}, negaBinary...)
	}

	return negaBinary
}

// ConvBinaryToDecimal converts binary (base 2) in 0-1 array format to decimal (base 10)
func ConvBinaryToDecimal(value []int) int {
	decimal := 0

	for i, b := range value {
		power := len(value) - 1 - i
		decimal += int(math.Pow(2, float64(power))) * b
	}

	return decimal
}

// ConvNegaBinaryToDecimal converts negative binary (base -2) in 0-1 array format to decimal (base 10)format
func ConvNegaBinaryToDecimal(value []int) int {
	decimal := 0

	for i, b := range value {
		power := len(value) - 1 - i
		decimal += int(math.Pow(-2, float64(power))) * b
	}

	return decimal
}

// AddBinary performs binary addition of a and b, both in 0-1 array format
func AddBinary(a, b []int) (result []int) {
	a, b = MakeSameLength(a, b)

	var digit, carry int
	for i := len(a) - 1; i >= 0; i-- {
		digit, carry = AddBinaryDigit(a[i], b[i], carry)
		result = append([]int{digit}, result...)
	}

	if carry == 1 {
		result = append([]int{carry}, result...)
	}

	return result
}

// MakeSameLength pads either a or b with leading zeros, so both have the same length
func MakeSameLength(a, b []int) ([]int, []int) {
	if len(b) > len(a) {
		a = PadLeftWithZero(a, len(b)-len(a))
	} else if len(a) > len(b) {
		b = PadLeftWithZero(b, len(a)-len(b))
	}

	return a, b
}

// PadLeftWithZero adds leading zeros prefix to a, with zeroCount length
func PadLeftWithZero(a []int, zeroCount int) []int {
	prefix := make([]int, zeroCount)

	return append(prefix, a...)
}

// AddBinaryDigit performs single binary digit addition of a and b, with optional carry
func AddBinaryDigit(a, b, carry int) (result int, carryOut int) {
	result = a + b + carry
	if result == 2 {
		result = 0
		carryOut = 1
	} else if result == 3 {
		result = 1
		carryOut = 1
	}

	return result, carryOut
}

// Factorial returns the factorial of a, e.g. 6! = 6 * 5 * 4 * 3 * 2 * 1 = 720
// SLOW!!!
func Factorial(a int) int {
	if a == 0 {
		return 1
	}

	return a * Factorial(a-1)
}

// BinomialCoefficient returns number of possible k combinations from n items
// SLOW!!!
func BinomialCoefficient(n, k int) int {
	if n < k {
		return 0
	}

	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

// Mean calculates the average of the array a
func Mean(a []int) float64 {
	var sum float64
	for _, v := range a {
		sum += float64(v)
	}

	return sum / float64(len(a))
}

// Median calculates the middle value of the array a
func Median(a []int) float64 {
	sort.Ints(a)

	if len(a)%2 == 1 {
		midIndex := (len(a) + 1) / 2

		return float64(a[midIndex])
	} else {
		midIndex1 := len(a)/2 - 1
		midIndex2 := midIndex1 + 1
		sum := a[midIndex1] + a[midIndex2]

		return float64(sum) / 2
	}
}
