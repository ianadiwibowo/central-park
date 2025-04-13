from collections import defaultdict
from typing import List


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graphs = defaultdict(list)
        for origin_node, destination_node, delay_time in times:
            graphs[origin_node].append((destination_node, delay_time))

        distances = {}
        for node in range(1, n+1):
            distances[node] = float("inf")

        def propagate(node, time):
            if time >= distances[node]:
                return
            distances[node] = time
            for neighbor_node, weight in graphs[node]:
                propagate(neighbor_node, time + weight)

        propagate(k, 0)
        shortest_time_to_visit_all = max(distances.values())
        if shortest_time_to_visit_all == float("inf"):
            return -1
        return shortest_time_to_visit_all


s = Solution()
print(s.networkDelayTime([[1, 2, 1], [2, 3, 1], [1, 4, 4], [3, 4, 1]], 4, 1))  # 3
print(s.networkDelayTime([[1, 2, 1], [2, 3, 1]], 3, 2))  # -1
