/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil || root == p || root == q {
    return root    
  }
  pl := lowestCommonAncestor(root.Left, p, q)
  pr := lowestCommonAncestor(root.Right, p, q)
  if pl != nil && pr != nil {
    return root
  }
  if pl == nil {
    return pr
  } else {
    return pl
  }
}