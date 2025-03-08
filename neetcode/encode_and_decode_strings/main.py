from typing import List
import io

class Solution:

    def encode(self, strs: List[str]) -> str:
        count = len(strs)
        if count == 0:
            return ""

        strs_element_lengths_in_char = []
        for i in strs:
            length_in_char = chr(len(i))
            strs_element_lengths_in_char.append(length_in_char)

        result = io.StringIO()
        for i, v in enumerate(strs_element_lengths_in_char):
            result.write(v)
            result.write(strs[i])

        return result.getvalue()

    def decode(self, s: str) -> List[str]:
        if s == "":
            return []

        index = 0
        result = []
        while index < len(s):
            element_length = ord(s[index])
            index += 1
            element = s[index:index+element_length]
            index += element_length
            result.append(element)

        return result

solution = Solution()
input = [""]
print(input)
x = solution.encode(input)
print(x)
y = solution.decode(x)
print(y)
