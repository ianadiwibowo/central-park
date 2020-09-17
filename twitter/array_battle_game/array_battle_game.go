package main

import "fmt"

func main() {
	// Rule: We pick leftmost/rightmost number, enemy does the same, in turns
	// Whomever collects the biggest sum wins
	arr := []int{1, 2, 3, 4, 5, 6}
	moves := [][]int{}

	findAllMovements(arr, []int{}, &moves)
	// fmt.Println("moves:", moves)

	fmt.Println(findFirstMoveWithTheHighestScore(moves))
}

func findAllMovements(arr []int, list []int, moves *[][]int) {
	if len(arr) == 0 {
		*moves = append(*moves, list)
		return
	}

	// processLeft()
	left := arr[0]
	leftList := make([]int, len(list)+1)
	copy(leftList, append(list, left))
	leftArr := arr[1:]
	// fmt.Println("Pick Left:", left, "- Remaining:", leftArr)
	findAllMovements(leftArr, leftList, moves)

	if len(leftArr) == 0 {
		return
	}

	// processRight()
	right := arr[len(arr)-1]
	rightList := make([]int, len(list)+1)
	copy(rightList, append(list, right))
	rightArr := arr[:len(arr)-1]
	// fmt.Println("Pick Right:", right, "- Remaining:", rightArr)
	findAllMovements(rightArr, rightList, moves)
}

func findFirstMoveWithTheHighestScore(moves [][]int) int {
	diffs := make([]int, len(moves))
	for i, m := range moves {
		// fmt.Printf("%v - ", m)
		moveCount := len(m)
		ourOddSum := 0
		for i := 0; i < moveCount; i += 2 {
			ourOddSum += m[i]
		}
		enemyEvenSum := 0
		for i := 1; i < moveCount; i += 2 {
			enemyEvenSum += m[i]
		}
		diff := ourOddSum - enemyEvenSum
		diffs[i] = diff
		// fmt.Printf("%v, %v, %v\n", ourOddSum, enemyEvenSum, diff)
		// fmt.Printf("Our Sum: %v, Enemy Sum: %v\n", ourOddSum, enemyEvenSum)
	}

	// findHighestDiff()
	highestDiff := -9999
	highestIndex := -1
	for i, d := range diffs {
		if d > highestDiff {
			highestDiff = d
			highestIndex = i
		}
	}
	// fmt.Printf("%v %v %v", highestDiff, highestIndex, moves[highestIndex][0])
	return moves[highestIndex][0]
}
