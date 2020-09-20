package sort

func QuickSort(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)

	quickSort(b, 0, len(b)-1)
	return b
}

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(a, left, right)
	quickSort(a, left, pivot-1)
	quickSort(a, pivot+1, right)
}

func partition(a []int, left, right int) int {
	pivot := a[right] // One of pivot technique, make pivot from last element
	for i := left; i < right; i++ {
		if a[i] < pivot {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	return left
}
