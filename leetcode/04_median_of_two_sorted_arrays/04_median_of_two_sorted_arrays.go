package main

// 4. Median of Two Sorted Arrays (Hard)
// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedNums := []int{}
	nums1Length := len(nums1)
	nums2Length := len(nums2)
	totalMergedArrayLength := nums1Length + nums2Length
	medianIndex := totalMergedArrayLength / 2
	nums1Index := 0
	nums2Index := 0

	for i := 0; i <= medianIndex; i++ {
		if nums1Index < nums1Length && nums2Index < nums2Length { // When both nums still have remaining elements
			if nums1[nums1Index] <= nums2[nums2Index] {
				mergedNums = append(mergedNums, nums1[nums1Index])
				nums1Index += 1
			} else {
				mergedNums = append(mergedNums, nums2[nums2Index])
				nums2Index += 1
			}
		} else if nums1Index < nums1Length { // When only nums1 remaining
			mergedNums = append(mergedNums, nums1[nums1Index])
			nums1Index += 1
		} else { // When only nums2 remaining
			mergedNums = append(mergedNums, nums2[nums2Index])
			nums2Index += 1
		}
	}

	if totalMergedArrayLength%2 != 0 { // Odd-length case
		return float64(mergedNums[len(mergedNums)-1])
	} else { // Even-length case
		return (float64(mergedNums[len(mergedNums)-1]) + float64(mergedNums[len(mergedNums)-2])) / 2
	}
}
