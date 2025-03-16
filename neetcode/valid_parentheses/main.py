from collections import deque


class Solution:
    def isValid(self, s: str) -> bool:
        stack = deque()
        for char in s:
            if char == "{" or char == "(" or char == "[":
                stack.append(char)
                continue

            if len(stack) == 0:
                return False

            previous_char = stack.pop()
            if char == "}" and previous_char != "{":
                return False
            if char == ")" and previous_char != "(":
                return False
            if char == "]" and previous_char != "[":
                return False

        return len(stack) == 0


solution = Solution()
print(solution.isValid("[]"))  # True
print(solution.isValid("([{}])"))  # True
print(solution.isValid("[(])"))  # False
print(solution.isValid("{}"))  # True
print(solution.isValid("("))  # False
print(solution.isValid(")"))  # False
