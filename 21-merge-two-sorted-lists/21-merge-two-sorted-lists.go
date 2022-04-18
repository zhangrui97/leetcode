/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  ans := &ListNode{}
  p := ans
  for list1 != nil && list2 != nil {
    if list1.Val < list2.Val {
      p.Next = list1
      p = p.Next
      list1 = list1.Next
    } else {
      p.Next = list2
      p = p.Next
      list2 = list2.Next
    }
  }
  if list1 != nil {
    p.Next = list1
  } else if list2 != nil {
    p.Next = list2
  }
  return ans.Next
}