package sort

// BubbleSort sorts the array a using Bubble Sort algorithm
func BubbleSort(a []int) (sorted []int) {
	sorted = make([]int, len(a))
	copy(sorted, a)

	n := len(sorted)
	for {
		swapped := false

		for i := 1; i < n; i++ {
			if sorted[i-1] > sorted[i] {
				temp := sorted[i-1]
				sorted[i-1] = sorted[i]
				sorted[i] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
		n--
	}

	return sorted
}
