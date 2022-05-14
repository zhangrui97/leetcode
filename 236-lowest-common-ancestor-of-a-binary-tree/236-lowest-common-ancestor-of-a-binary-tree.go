/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func contains(t *TreeNode, n *TreeNode)bool {
  if t == nil {
    return false
  }
  if t == n || (t.Left == n || t.Right == n) {
    return true
  }
  if t.Left == nil && t.Right == nil {
    return false
  }
  return contains(t.Left, n) || contains(t.Right, n)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == p || root == q {
    return root    
  }
  pp := contains(root.Left, p)
  pq := contains(root.Right, q)
  if (pp && pq) || (!pp && !pq) {
    return root
  }
  if pp {
    return lowestCommonAncestor(root.Left, p, q)
  } else {
    return lowestCommonAncestor(root.Right, p, q)
  }
}