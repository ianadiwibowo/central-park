from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = nums[0]
        current_sum = nums[0]
        for i in range(1, len(nums)):
            # For each next element, try extending the subarray first
            extended_subarray_sum = current_sum + nums[i]

            # But, if turns out the current element is bigger than the existing subarray sum
            # ditch the existing subarray, and start a new one starting this element
            if nums[i] > extended_subarray_sum:
                current_sum = nums[i]
            # But, if not, then keep extending the subarray added with the current element
            else:
                current_sum = extended_subarray_sum

            # Update the max_sum whenever the current_sum is bigger
            if current_sum > max_sum:
                max_sum = current_sum
        return max_sum
