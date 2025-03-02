from collections import defaultdict
from typing import List

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = defaultdict(int)
        for i in nums:
            freq[i] += 1

        buckets = [None] * (len(nums) + 1)
        for key, value in freq.items():
            if buckets[value] is None:
                buckets[value] = []
            buckets[value].append(key)

        result = []
        for i in range(len(buckets)-1, -1, -1):
            if buckets[i] is None:
                continue

            if k == len(buckets[i]):
                result += buckets[i]
                return result
            elif k < len(buckets[i]):
                result += buckets[i][0:k]
                return result
            else: # k > len(buckets[i]):
                result += buckets[i]
                k -= len(buckets[i])

        return result

nums = [-1000, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 5]
# nums = [7,7]
k = 6

solution = Solution()
print(solution.topKFrequent(nums, k))
