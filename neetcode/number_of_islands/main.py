from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        def explore_the_entire_island(r: int, c: int):
            if r < 0 or r >= len(grid) or c < 0 or c >= len(grid[r]):
                return
            if grid[r][c] == "0":
                return
            grid[r][c] = "0"  # Mark as visited
            explore_the_entire_island(r-1, c)
            explore_the_entire_island(r+1, c)
            explore_the_entire_island(r, c-1)
            explore_the_entire_island(r, c+1)

        islands_count = 0
        for r in range(len(grid)):
            for c in range(len(grid[r])):
                if grid[r][c] == "1":
                    islands_count += 1
                    explore_the_entire_island(r, c)

        return islands_count


s = Solution()

grid = [
    ["0", "1", "1", "1", "0"],
    ["0", "1", "0", "1", "0"],
    ["1", "1", "0", "0", "0"],
    ["0", "0", "0", "0", "0"]
]
print(s.numIslands(grid))  # 1

grid = [
    ["1", "1", "0", "0", "1"],
    ["1", "1", "0", "0", "1"],
    ["0", "0", "1", "0", "0"],
    ["0", "0", "0", "1", "1"]
]
print(s.numIslands(grid))  # 4

grid = [
    ["0", "0", "0", "0", "0"],
    ["0", "0", "0", "0", "0"],
    ["0", "0", "0", "0", "0"],
    ["0", "0", "0", "0", "0"]
]
print(s.numIslands(grid))  # 0

grid = [
    ["1", "1", "1", "1", "1"],
    ["1", "1", "1", "1", "1"],
    ["1", "1", "1", "1", "1"],
    ["1", "1", "1", "1", "1"],
]
print(s.numIslands(grid))  # 1

grid = [
    ["1", "0", "1", "0", "1"],
    ["0", "1", "0", "1", "0"],
    ["1", "0", "1", "0", "1"],
    ["0", "1", "0", "1", "0"],
]
print(s.numIslands(grid))  # 10

grid = [
    ["1", "0", "1", "0", "1"],
]
print(s.numIslands(grid))  # 3

grid = [
    ["0", "0", "0", "0", "0"],
]
print(s.numIslands(grid))  # 0

grid = [
    ["1", "1", "1", "1", "1"],
]
print(s.numIslands(grid))  # 1

grid = [
    ["1"],
    ["0"],
    ["1"],
    ["0"],
    ["1"],
]
print(s.numIslands(grid))  # 3

grid = [
    ["1"],
    ["1"],
    ["1"],
    ["1"],
    ["1"],
]
print(s.numIslands(grid))  # 1

grid = [
    ["0"],
    ["0"],
    ["0"],
    ["0"],
    ["0"],
]
print(s.numIslands(grid))  # 0

grid = [
    ["0"],
]
print(s.numIslands(grid))  # 0

grid = [
    ["1"],
]
print(s.numIslands(grid))  # 1
