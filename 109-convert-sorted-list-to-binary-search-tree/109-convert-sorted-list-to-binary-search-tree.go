/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
  if head == nil { return nil }
  if head.Next == nil { return &TreeNode{head.Val, nil, nil} } 
  pre, sp, fp := head, head, head
  for fp != nil && fp.Next != nil {
    pre = sp
    sp = sp.Next
    fp = fp.Next.Next
  }
  pre.Next = nil
  return &TreeNode{sp.Val, sortedListToBST(head), sortedListToBST(sp.Next)}
}