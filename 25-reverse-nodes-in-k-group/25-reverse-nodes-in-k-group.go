var tail *ListNode = nil
var fullLength = false
func reverse(head *ListNode, k int) *ListNode {
  if head == nil { return head }
  if head.Next == nil || k == 1 {
    tail = head.Next
    fullLength = (k == 1)
    return head
  }  
  reversed := reverse(head.Next, k - 1)
  head.Next.Next = head
  head.Next = tail
  return reversed
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
  if k == 1 || head == nil {
    return head
  }
  reversed := reverse(head, k)
  if fullLength {
    head.Next = reverseKGroup(tail, k)
  } else {
    reversed = reverse(reversed, k)
  }
  return reversed
}