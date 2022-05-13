/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSym(p, q *TreeNode) bool {
  if p == nil && q == nil {
    return true
  }
  if p == nil || q == nil {
    return false
  }
  return p.Val == q.Val && isSym(p.Left, q.Right) && isSym(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
  return root == nil || isSym(root.Left, root.Right)    
}