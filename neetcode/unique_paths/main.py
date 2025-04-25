class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        cache = [[0] * n for _ in range(m)]
        for col in range(n):
            cache[0][col] = 1
        for row in range(m):
            cache[row][0] = 1
        for row in range(1, m):
            for col in range(1, n):
                cache[row][col] = cache[row-1][col] + cache[row][col-1]
        return cache[m-1][n-1]
