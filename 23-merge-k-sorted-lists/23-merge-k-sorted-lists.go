func merge(l1, l2 *ListNode) *ListNode {
  ans := &ListNode{}
  p := ans
  for l1 != nil && l2 != nil {
    if l1.Val < l2.Val {
      p.Next = l1
      l1 = l1.Next
    } else {
      p.Next = l2
      l2 = l2.Next
    }
    p = p.Next
  }
  if l1 != nil {
    p.Next = l1
  }
  if l2 != nil {
    p.Next = l2
  }
  return ans.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
  l := len(lists)
  switch l {
  case 0: return nil
  case 1: return lists[0]
  default:
    lists[l-2] = merge(lists[l-2], lists[l-1])
    return mergeKLists(lists[:l-1])
  }
}