# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val: int, next: Optional[ListNode]):
        self.val = val
        self.next = next

class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        position = 0
        length = 0
        curr = head

        while curr:
            curr = curr.next
            length += 1
        
        curr = head
        while position < int(length / 2) and curr:
            if position == int(length / 2) - 1:
                if curr.next and curr.next.next:
                    temp = curr.next.next
                    curr.next = temp
                else:
                    curr.next = None
                break
            curr = curr.next
            position += 1

        print(length)

        return head