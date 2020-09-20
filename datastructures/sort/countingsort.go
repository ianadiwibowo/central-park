package sort

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
