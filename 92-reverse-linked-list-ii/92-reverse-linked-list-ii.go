func reverse(head *ListNode, count int) *ListNode {
  if count == 0 { return head }  
  next := head.Next
  last := reverse(next, count-1)
  head.Next = next.Next
  next.Next = head
  return last
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
  if left >= right { return head }
  if left == 1 { return reverse(head, right - left) }
  pre, p := head, head
  for i := 0; i < left-1; i++ {
    pre, p = p, p.Next
  }
  pre.Next = reverse(p, right - left)  
  return head
}