package permutation

func HeapPermutation(a []int) [][]int {
	results := [][]int{}

	var heapPermutation func(a []int, size int)
	heapPermutation = func(a []int, size int) {
		if size == 1 {
			b := make([]int, len(a))
			copy(b, a)
			results = append(results, b)
			return
		}

		for i := 0; i < size; i++ {
			heapPermutation(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}

	heapPermutation(a, len(a))
	return results
}
