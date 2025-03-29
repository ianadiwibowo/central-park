from typing import Optional
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        self.invertChildren(root)
        return root

    def invertChildren(self, node: Optional[TreeNode]):
        if node is None:
            return

        temp_node = node.left
        node.left = node.right
        node.right = temp_node

        self.invertChildren(node.left)
        self.invertChildren(node.right)

    # Bread-First Traversal
    def traverse_level_order(self, root: Optional[TreeNode]) -> list[int]:
        if root is None:
            return []
        queue = deque()
        queue.append(root)
        result = []
        while queue:
            current_node = queue.popleft()
            result.append(current_node.val)
            if current_node.left:
                queue.append(current_node.left)
            if current_node.right:
                queue.append(current_node.right)
        return result


solution = Solution()

# Full balanced tree
# Input:
#    1
#  2   3
# 4 5 6 7
# Output:
#    1
#  3   2
# 7 6 5 4
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.left.left = TreeNode(4)
root.left.right = TreeNode(5)
root.right.left = TreeNode(6)
root.right.right = TreeNode(7)

result = solution.invertTree(root)
print(solution.traverse_level_order(result))  # [1, 3, 2, 7, 6, 5, 4]

print("------")

# Empty tree
result = solution.invertTree(None)
print(solution.traverse_level_order(result))  # []

print("------")

# Root only without left
root = TreeNode(1)
root.right = TreeNode(3)

result = solution.invertTree(root)
print(solution.traverse_level_order(result))  # [1, 3]
print(root.left.val == 3)
print(root.right is None)

print("------")

# Root only without right
root = TreeNode(1)
root.left = TreeNode(2)

result = solution.invertTree(root)
print(solution.traverse_level_order(result))  # [1, 2]
print(root.right.val == 2)
print(root.left is None)
