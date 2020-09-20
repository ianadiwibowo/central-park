package sort

func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		currentValue := a[i]
		j := i

		for j > 0 && currentValue < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}

	return a
}
