/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
  var traverse func(t *TreeNode, v int)int
  traverse = func(t *TreeNode, v int)int {
    if t == nil { return v }
    t.Val += traverse(t.Right, v)
    return traverse(t.Left, t.Val)
  }
  traverse(root, 0)
  return root
}