package sort

import "fmt"

func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		fmt.Println("i:", i, "a[i]:", a[i], "—", a)
		currentValue := a[i]
		j := i

		for j > 0 && currentValue < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}

	return a
}
