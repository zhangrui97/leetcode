/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSame(t, sub *TreeNode) bool {
  if t == nil && sub == nil { return true }
  if t == nil || sub == nil { return false }
  return t.Val == sub.Val && isSame(t.Left, sub.Left) && isSame(t.Right, sub.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil { return false }
  if root.Val == subRoot.Val && isSame(root, subRoot) { return true }
  return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}