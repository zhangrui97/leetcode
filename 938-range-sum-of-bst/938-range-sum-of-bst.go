/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
  if root == nil { return 0 }  
  if root.Val < low { return rangeSumBST(root.Right, low, high) }
  if root.Val > high { return rangeSumBST(root.Left, low, high) }
  val := root.Val
  return val + rangeSumBST(root.Left, low, val-1) + rangeSumBST(root.Right, val+1, high)
}