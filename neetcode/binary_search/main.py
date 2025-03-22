from typing import List
import math


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums) - 1

        while left <= right:
            if left == right:
                mid = left
            else:
                mid = left + math.ceil((right - left) / 2)

            if target == nums[mid]:
                return mid
            elif target < nums[mid]:
                right = mid - 1
            elif target > nums[mid]:
                left = mid + 1

        return -1

    def searchV2(self, nums: List[int], target: int) -> int:
        return self.binarySearch(0, len(nums)-1, nums, target)

    def binarySearch(self, left: int, right: int, nums: List[int], target: int) -> int:
        if left > right:
            return -1
        mid = left + ((right - left) // 2)

        if target == nums[mid]:
            return mid
        elif target < nums[mid]:
            return self.binarySearch(left, mid-1, nums, target)
        else:  # target > nums[mid]:
            return self.binarySearch(mid+1, right, nums, target)


solution = Solution()

print(solution.search([5], 5) == 0)
print(solution.search([5], 10) == -1)
print(solution.search([1, 2, 3], 2) == 1)
print(solution.search([1, 5, 10], 6) == -1)
print(solution.search([-1, 0, 2, 4, 6, 8], 4) == 3)
print(solution.search([-1, 0, 2, 4, 6, 8], 3) == -1)
print(solution.search([-1, 0, 2, 4, 6, 8], 2) == 2)
print(solution.search([0, 1, 2, 3, 4, 5], 5) == 5)
print(solution.search([0, 1, 2, 3, 4, 5], 6) == -1)

print(solution.searchV2([5], 5) == 0)
print(solution.searchV2([5], 10) == -1)
print(solution.searchV2([1, 2, 3], 2) == 1)
print(solution.searchV2([1, 5, 10], 6) == -1)
print(solution.searchV2([-1, 0, 2, 4, 6, 8], 4) == 3)
print(solution.searchV2([-1, 0, 2, 4, 6, 8], 3) == -1)
print(solution.searchV2([-1, 0, 2, 4, 6, 8], 2) == 2)
print(solution.searchV2([0, 1, 2, 3, 4, 5], 5) == 5)
print(solution.searchV2([0, 1, 2, 3, 4, 5], 6) == -1)
