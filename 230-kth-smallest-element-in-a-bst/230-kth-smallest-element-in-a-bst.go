/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
  var traverse func(t *TreeNode)*TreeNode
  traverse = func(t *TreeNode)*TreeNode {
    if t == nil { return nil }
    result := traverse(t.Left) 
    if result != nil {
      return result
    }
    k--
    if k == 0 {
      return t
    }
    return traverse(t.Right)
  }
  return traverse(root).Val
}