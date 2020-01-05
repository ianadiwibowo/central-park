package array

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

// Equal checks whether two slices a and b are identical (true) or not (false)
// DEPRECATED, use cmp.Equal instead
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
