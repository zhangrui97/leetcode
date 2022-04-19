/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
  prob, ans := head, head
  for prob != nil && prob.Next != nil {
    prob = prob.Next.Next
    ans = ans.Next
  }
  return ans
}