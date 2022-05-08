/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
  if root == nil {
    return
  } 
  left, right := root.Left, root.Right
  flatten(left)
  flatten(right)
  root.Left, root.Right = nil, left
  p := root
  for p.Right != nil {
    p = p.Right
  }
  p.Right = right
}