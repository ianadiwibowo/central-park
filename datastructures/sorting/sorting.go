package sorting

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

// CountingSort sorts the array a using Counting Sort algorithm
func CountingSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	counter := make([]int, len(a))
	for _, v := range a {
		counter[v]++
	}

	for i := 1; i < len(counter); i++ {
		counter[i] = counter[i] + counter[i-1]
	}

	sorted := make([]int, len(a))
	for _, v := range a {
		sorted[counter[v]-1] = v
		counter[v]--
	}

	return sorted
}
