package main

import "fmt"

// https://www.hackerrank.com/challenges/new-year-chaos/problem
// Complete the minimumBribes function below.
func minimumBribes(q []int32) string {
	bribesCount := make(map[int32]int)
	n := len(q)

	for {
		bribed := false
		for i := 1; i < n; i++ {
			if q[i-1] > q[i] {
				bribesCount[q[i-1]] += 1
				if bribesCount[q[i-1]] > 2 {
					return "Too chaotic"
				}

				temp := q[i-1]
				q[i-1] = q[i]
				q[i] = temp

				bribed = true
			}
		}
		n--

		if !bribed {
			break
		}
	}

	totalBribes := 0
	for _, value := range bribesCount {
		totalBribes += value
	}
	return fmt.Sprintf("%v", totalBribes)
}

func main() {
	q_cases := [][]int32{
		{2, 1, 5, 3, 4},
		{2, 5, 1, 3, 4},
		{5, 1, 2, 3, 7, 8, 6, 4},
		{1, 2, 5, 3, 7, 8, 6, 4},
		{1, 2, 5, 3, 4, 7, 8, 6},
	}
	expectations := []string{
		"3",
		"Too chaotic",
		"Too chaotic",
		"7",
		"4",
	}

	for i, v := range q_cases {
		result := minimumBribes(v)
		fmt.Printf("Passed: %v. Test: %v. Expected: %v. Got: %v.\n",
			result == expectations[i],
			v,
			expectations[i],
			result,
		)
	}
}
