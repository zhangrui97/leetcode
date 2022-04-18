func removeEmpty(lists []*ListNode) []*ListNode {
  slow := 0
  for _, v := range lists {
    if v != nil {
      lists[slow] = v
      slow++
    }
  }
  return lists[:slow]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
  lists = removeEmpty(lists)
  switch len(lists) {
  case 0: return nil
    case 1: return lists[0]
  default:
    ans := &ListNode{}
    p := ans
    for len(lists) > 1 {
      sm := 0
      for i := 1; i < len(lists); i++ {
        if lists[i].Val < lists[sm].Val {
          sm = i
        }
      }
      p.Next = lists[sm]
      p = p.Next
      if lists[sm].Next == nil {
        lists[len(lists)-1], lists[sm] = lists[sm], lists[len(lists)-1]
        lists = lists[:len(lists)-1]
      } else {
        lists[sm] = lists[sm].Next
      }
    }
    p.Next = lists[0]
    return ans.Next
  }
}