class Solution:
    def climbStairs(self, n: int) -> int:
        cache = [0, 1, 2]
        if n <= 2:
            return cache[n]
        for i in range(3, n + 1):
            cache.append(cache[i - 1] + cache[i - 2])
        return cache[-1]
        # 1 there's 1
        # 1
        # 2 there's 2
        # 1 1
        # 2
        # 3 there's 3
        # 1 1 1
        # 2 1
        # 1 2
        # 4 there's 5
        # 1 1 1 1
        # 2 1 1
        # 1 2 1
        # 1 1 2
        # 2 2
        # 5 there's 8
        # 1 1 1 1 1
        # 2 1 1 1
        # 1 2 1 1
        # 1 1 2 1
        # 1 1 1 2
        # 2 2 1
        # 2 1 2
        # 1 2 2
        # 6 there's 13
        # 1 1 1 1 1 1
        # 2 1 1 1 1
        # 1 2 1 1 1
        # 1 1 2 1 1
        # 1 1 1 2 1
        # 1 1 1 1 2
        # 2 2 1 1
        # 2 1 2 1
        # 2 1 1 2
        # 1 2 2 1
        # 1 2 1 2
        # 1 1 2 2
        # 2 2 2
