package main

import (
	"fmt"
	"strings"
	"time"
)

func NoCornerSquare(squareSize int) string {
	if squareSize <= 0 {
		return "Please insert a valid number"
	}

	firstAndLastRow := constructRow(squareSize, " ", "x")
	middleRow := constructRow(squareSize, "x", " ")

	var result strings.Builder

	result.WriteString(firstAndLastRow)
	for row := 0; row < squareSize; row++ {
		result.WriteString(middleRow)
	}
	result.WriteString(firstAndLastRow)

	return removeLastNewline(result.String())
}

func constructRow(n int, char1, char2 string) (result string) {
	var res strings.Builder

	res.WriteString(char1)
	res.WriteString(constructColumn(n, char2))
	res.WriteString(char1)
	res.WriteString("\n")

	return res.String()
}

func constructColumn(n int, char string) string {
	return strings.Repeat(char, n)
}

func removeLastNewline(s string) string {
	return s[:len(s)-1]
}

func main() {
	for squareSize := 0; squareSize <= 50000; squareSize += 2000 {
		measure(squareSize)
	}
}

func measure(squareSize int) {
	defer timeTrack(time.Now(), fmt.Sprintf("%v", squareSize))
	_ = NoCornerSquare(squareSize)
}

func timeTrack(start time.Time, identifier string) {
	elapsed := time.Since(start)
	fmt.Printf("%v, %v\n", identifier, elapsed.Milliseconds())
}
