package search

func BinarySearch(a []int, value int) (index int) {
	left := 0
	right := len(a) - 1

	for left <= right {
		mid := (left + right) / 2

		if a[mid] < value {
			left = mid + 1
		} else if a[mid] > value {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
