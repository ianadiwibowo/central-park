package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	a.Equal(t, "A", convert("A", 1))
	a.Equal(t, "ACBD", convert("ABCD", 2))
	a.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	a.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
}
