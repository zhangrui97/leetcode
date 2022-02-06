# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
  def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
    p1 = l1
    p2 = l2
    result = ListNode()
    pre = result
    p3 = result
    carry = False
    while p1 or p2:
      sum = 1 if carry else 0
      if p1: 
        sum += p1.val
        p1 = p1.next
      if p2: 
        sum += p2.val
        p2 = p2.next
      carry = True if sum >= 10 else False
      p3.val = sum % 10
      p3.next = ListNode()
      pre = p3
      p3 = p3.next
    if carry: 
      p3.val = 1
    else:
      pre.next = None
    return result