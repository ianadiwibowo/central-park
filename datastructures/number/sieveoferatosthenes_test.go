package number_test

import (
	"testing"

	"github.com/ianadiwibowo/central-park/datastructures/number"
	"github.com/stretchr/testify/assert"
)

func TestSieveOfEratosthenes_WithLimit50(t *testing.T) {
	primeNumbers50 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47}
	assert.Equal(t, primeNumbers50, number.SieveOfEratosthenes(50))
}

func TestSieveOfEratosthenes_WithLimit100(t *testing.T) {
	primeNumbers100 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	assert.Equal(t, primeNumbers100, number.SieveOfEratosthenes(100))
}
