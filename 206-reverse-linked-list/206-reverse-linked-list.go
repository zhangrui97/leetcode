/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  tail := head.Next
  newHead := reverseList(tail)
  head.Next = nil
  tail.Next = head
  return newHead
}