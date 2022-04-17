/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil { return head }    
  slow, fast := head, head
  for fast != nil {
    if fast.Val != slow.Val {
      slow.Next = fast
      slow = fast
    }
    fast = fast.Next;
  }
  slow.Next = nil
  return head
}