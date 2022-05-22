/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
  if root == nil { return nil }
  l := increasingBST(root.Left)
  root.Left = nil
  root.Right = increasingBST(root.Right)
  if l == nil { 
    return root
  } else {
    r := l
    for r.Right != nil {
      r = r.Right
    }
    r.Right = root
    return l
  }
}