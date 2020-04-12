package main

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestNoCornerSquare_GivenInvalidInput_ReturnsErrorMessage(t *testing.T) {
	expectation := "Please insert a valid number"
	assert(t, 0, expectation)
	assert(t, -5, expectation)
}

func TestNoCornerSquare_GivenValidInput_DrawsCorrectly(t *testing.T) {
	assert(t, 1, " x \nx x\n x ")
	assert(t, 2, " xx \nx  x\nx  x\n xx ")
	assert(t, 3, " xxx \nx   x\nx   x\nx   x\n xxx ")
	assert(t, 4, " xxxx \nx    x\nx    x\nx    x\nx    x\n xxxx ")
	assert(t, 5, " xxxxx \nx     x\nx     x\nx     x\nx     x\nx     x\n xxxxx ")
}

func assert(t *testing.T, squareSize int, expectation string) bool {
	return testifyAssert.Equal(t, expectation, NoCornerSquare(squareSize))
}
