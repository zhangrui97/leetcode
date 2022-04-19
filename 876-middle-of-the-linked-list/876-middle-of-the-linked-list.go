/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
  prob, ans := head, head
  for prob != nil {
    prob = prob.Next
    if prob == nil { break }
    prob = prob.Next
    ans = ans.Next
  }
  return ans
}