class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.lower()
        front_index = 0
        back_index = len(s)-1

        while front_index <= back_index:
            if not s[front_index].isalnum():
                front_index += 1
                continue
            if not s[back_index].isalnum():
                back_index -= 1
                continue
            if s[front_index] != s[back_index]:
                return False
            front_index += 1
            back_index -= 1

        return True


solution = Solution()
print(solution.isPalindrome("Was it a car or a cat I saw?"))
print(solution.isPalindrome("tab a cat"))
print(solution.isPalindrome("X"))
print(solution.isPalindrome("-"))
print(solution.isPalindrome("aIBohphOBiA"))
print(solution.isPalindrome("a IB oh phOB---?iA"))
