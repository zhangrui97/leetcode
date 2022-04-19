/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  tail := head
  for i := 0; i < n; i++ {
    tail = tail.Next
  }
  var pre *ListNode
  ans := head
  for ; tail != nil; tail = tail.Next {
    pre = ans
    ans = ans.Next
  }
  if ans == head {
    return head.Next
  } else {
    pre.Next = ans.Next
    return head
  }
}