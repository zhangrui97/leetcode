/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
  if root == nil || (root.Left == nil && root.Right == nil) { return true }
  v := root.Val
  if root.Left != nil && root.Left.Val != v { return false }
  if root.Right != nil && root.Right.Val != v { return false }
  return isUnivalTree(root.Left) && isUnivalTree(root.Right) 
}