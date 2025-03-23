from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # O(n) complexity, with O(n) space
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        list = []
        current_node = head

        while current_node is not None:
            list.append(current_node.val)
            current_node = current_node.next

        new_head = None
        current_node = None
        prev_node = None

        for i in range(len(list) - 1, -1, -1):
            current_node = ListNode(list[i], None)
            if i == len(list) - 1:
                new_head = current_node
            else:
                prev_node.next = current_node
            prev_node = current_node

        return new_head

    # O(n) complexity, with O(1) space
    def reverseListV2(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Zero case
        if head is None:
            return None

        # 1-size case
        if head.next is None:
            return head

        # 2-size case
        if head.next is not None and head.next.next is None:
            new_head = head.next
            new_head.next = head
            head.next = None
            return new_head

        # 3+-size case
        node1 = head
        node2 = head.next
        node3 = head.next.next

        while node2 is not None:
            node2.next = node1

            if node1 is head:
                node1.next = None

            # self.printLinkedList(node2)

            node1 = node2
            node2 = node3
            if node3 is not None:
                node3 = node3.next

        return node1

    def printLinkedList(self, head: Optional[ListNode]):
        print("---")

        if head is None:
            print("None")
            return

        while head is not None:
            print(head.val)
            head = head.next

# a -> b -> c -> None

# node1 = a
# node2 = b
# node3 = c

# b -> a -> None

# node1 = b
# node2 = c
# node3 = None

# c -> b -> a -> None

# node1 = c
# node2 = None
# node3 = None


solution = Solution()

node1 = ListNode(1, None)
node2 = ListNode(2, None)
node3 = ListNode(3, None)
node4 = ListNode(4, None)
node1.next = node2
node2.next = node3
node3.next = node4
result1 = solution.reverseListV2(node1)
solution.printLinkedList(result1)


result2 = solution.reverseListV2(None)
solution.printLinkedList(result2)


node1 = ListNode(7, None)
node2 = ListNode(90, None)
node3 = ListNode(-5, None)
node4 = ListNode(17600, None)
node5 = ListNode(1, None)
node1.next = node2
node2.next = node3
node3.next = node4
node4.next = node5


result3 = solution.reverseListV2(node1)
solution.printLinkedList(result3)


node1 = ListNode(2, None)
node2 = ListNode(3, None)
node1.next = node2
result4 = solution.reverseListV2(node1)
solution.printLinkedList(result4)
