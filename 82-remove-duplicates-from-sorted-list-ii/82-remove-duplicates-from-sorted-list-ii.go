/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil || head.Next == nil { return head }
  var ans, next, slow, fast *ListNode = nil, nil, head, head.Next
  for fast != nil {
    if slow.Val != fast.Val {
      if slow.Next == fast  {
        if ans == nil {
          ans = slow
          next = ans
        } else {
          next.Next = slow
          next = slow
        }
      }
      slow = fast
      if fast.Next == nil {
        if ans == nil {
          return fast
        } else {
          next.Next = fast
          return ans
        }
      }
    }
    fast = fast.Next
  }
  
  if ans != nil {
    next.Next = nil
  }
  return ans
}