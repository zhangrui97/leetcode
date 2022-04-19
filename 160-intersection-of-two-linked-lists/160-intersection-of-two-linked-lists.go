/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
  pa, pb := headA, headB
  for pa != nil && pb != nil {
    pa = pa.Next
    pb = pb.Next
  }
  p, pl, ps := pa, headA, headB
  if pb != nil {
    p, pl, ps = pb, headB, headA
  }
  for p != nil {
    p = p.Next
    pl = pl.Next
  }
  for pl != ps {
    pl = pl.Next
    ps = ps.Next
  }
  return pl
}