from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        subsets_so_far = []

        def depth_first_search(i):
            if i >= len(nums):
                result.append(subsets_so_far.copy())
                return

            subsets_so_far.append(nums[i])
            depth_first_search(i + 1)

            subsets_so_far.pop()
            depth_first_search(i + 1)

        depth_first_search(0)
        return result


solution = Solution()

# => [[], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]]
print(solution.subsets([1, 2, 3]))

# => [[], [7]]
print(solution.subsets([7]))
