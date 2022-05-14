/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func min(a, b int) int {
  if a > b {
    return b
  } else {
    return a
  }
}
func minDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }
  if root.Left == nil && root.Right == nil {
    return 1
  }
  if root.Left == nil || root.Right == nil {
    if root.Left == nil {
      return minDepth(root.Right) + 1
    } else {
      return minDepth(root.Left) + 1
    }
  }
  return min(minDepth(root.Left), minDepth(root.Right)) + 1
}