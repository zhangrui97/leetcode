/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
  if root == nil { 
    return 0
  }  
  ld, rd := 1, 1
  p := root.Left
  for p != nil {
    ld++
    p = p.Left
  }
  p = root.Right
  for p != nil {
    rd++
    p = p.Right
  }
  if ld == rd {
    return 1 << ld - 1
  }
  return 1 + countNodes(root.Left) + countNodes(root.Right)
}