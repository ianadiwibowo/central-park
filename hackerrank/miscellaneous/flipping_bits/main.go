package main

import "math"

func FlippingBits(input int64) int64 {
	binary := convertDecimalToBinary(input)
	binary = padLeftWithZero(binary, 32-len(binary))
	binary = flip(binary)

	return convertBinaryToDecimal(binary)
}

func convertDecimalToBinary(decimal int64) (binary []int64) {
	if decimal == 0 {
		return []int64{0}
	}

	for decimal != 0 {
		binary = append([]int64{decimal % 2}, binary...)
		decimal = decimal / 2
	}

	return binary
}

func padLeftWithZero(binary []int64, zeroCount int) []int64 {
	prefix := make([]int64, zeroCount)

	return append(prefix, binary...)
}

func flip(binary []int64) []int64 {
	flippedBinary := make([]int64, len(binary))

	for i, v := range binary {
		flippedBinary[i] = flipBit(v)
	}

	return flippedBinary
}

func flipBit(bit int64) int64 {
	if bit == 0 {
		return 1
	} else {
		return 0
	}
}

func convertBinaryToDecimal(binary []int64) int64 {
	var decimal int64

	for i, b := range binary {
		power := len(binary) - 1 - i
		decimal += int64(math.Pow(2, float64(power))) * b
	}

	return decimal
}
