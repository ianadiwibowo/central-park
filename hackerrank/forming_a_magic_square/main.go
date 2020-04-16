package main

func FormingMagicSquare(square [][]int32) int32 {
	possible3x3MagicSquares := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	}
	cheapestCost := int32(10000000)
	for _, s := range possible3x3MagicSquares {
		cost := calculateCost(square, s)
		if cost < cheapestCost {
			cheapestCost = cost
		}
	}
	return cheapestCost
}

func calculateCost(before, after [][]int32) int32 {
	var totalCost int32
	for i, brow := range before {
		for j, b := range brow {
			cost := b - after[i][j]
			if cost < 0 {
				cost = -cost
			}
			totalCost += cost
		}
	}
	return totalCost
}

// func isMagicSquare(square [][]int32) bool {
// 	if !isDistinct(square) {
// 		return false
// 	}

// 	var magicConstant int32

// 	// // Diagonals
// 	// 00 11 22
// 	// 20 11 02
// 	var dia001122Sum int32
// 	for rc := 0; rc < 3; rc++ {
// 		dia001122Sum += square[rc][rc]
// 	}
// 	magicConstant = dia001122Sum
// 	// fmt.Println("Dia 001122:", dia001122Sum)

// 	var dia201102Sum int32
// 	for rc := 0; rc < 3; rc++ {
// 		dia201102Sum += square[2-rc][rc]
// 	}
// 	// fmt.Println("Dia 201102:", dia201102Sum)
// 	if dia201102Sum != magicConstant {
// 		return false
// 	}

// 	// Rows
// 	// 00 01 02
// 	// 10 11 12
// 	// 20 21 22
// 	for r := 0; r < 3; r++ {
// 		var rowSum int32
// 		for c := 0; c < 3; c++ {
// 			rowSum += square[r][c]
// 		}
// 		// fmt.Println("Row:", r, rowSum)
// 		if rowSum != magicConstant {
// 			return false
// 		}
// 	}

// 	// Columns
// 	// 00 10 20
// 	// 01 11 21
// 	// 02 12 22
// 	for c := 0; c < 3; c++ {
// 		var colSum int32
// 		for r := 0; r < 3; r++ {
// 			colSum += square[r][c]
// 		}
// 		// fmt.Println("Col", c, colSum)
// 		if colSum != magicConstant {
// 			return false
// 		}
// 	}

// 	return true
// }

// func isDistinct(square [][]int32) bool {
// 	counter := make(map[int32]int)
// 	for _, row := range square {
// 		for _, col := range row {
// 			if col < 1 || col > 9 {
// 				return false
// 			}
// 			counter[col]++
// 			if counter[col] > 1 {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }
