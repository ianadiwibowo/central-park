class TreeNode:

    def __init__(self, val: str = "", children: dict = None, end_of_word: bool = False):
        self.val = val

        if children is None:
            self.children = {}
        else:
            self.children = children

        self.end_of_word = end_of_word


class PrefixTree:

    def __init__(self):
        self.root = TreeNode()

    def insert(self, word: str) -> None:
        current_node = self.root
        for char in word:
            if char not in current_node.children:
                current_node.children[char] = TreeNode(char)
            current_node = current_node.children[char]
        current_node.end_of_word = True

    def search(self, word: str) -> bool:
        current_node = self.root
        for char in word:
            if char not in current_node.children:
                return False
            current_node = current_node.children[char]
        if not current_node.end_of_word:
            return False
        return True

    def startsWith(self, prefix: str) -> bool:
        current_node = self.root
        for char in prefix:
            if char not in current_node.children:
                return False
            current_node = current_node.children[char]
        return True


pt = PrefixTree()
print(pt.insert("apple"))  # None
print(pt.search("apple"))  # True
print(pt.search("app"))  # False
print(pt.search("porcupine"))  # False
print(pt.insert("apps"))  # None
print(pt.search("apple"))  # True
print(pt.search("apps"))  # True
print(pt.search("application"))  # False
print(pt.startsWith("ap"))  # True
print(pt.startsWith("appl"))  # True
print(pt.startsWith("apple"))  # True
