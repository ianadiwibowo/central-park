from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy_index = 0
        profit = 0

        for sell_index in range(1, len(prices)):
            potential_profit = prices[sell_index] - prices[buy_index]

            if potential_profit <= 0 and sell_index < len(prices):
                buy_index = sell_index
                continue

            if potential_profit > profit:
                profit = potential_profit

        return profit


solution = Solution()
print(solution.maxProfit([10, 1, 5, 6, 7, 1])) # 6
print(solution.maxProfit([10, 8, 7, 5, 2])) # 0
print(solution.maxProfit([10, 1, 5, 6, 0, 7, 1]))  # 7
print(solution.maxProfit([10, 1, 5, 6, 0, 1, 1]))  # 5
print(solution.maxProfit([10, 1, 5, 6, 0, 1, 1, 10]))  # 10
