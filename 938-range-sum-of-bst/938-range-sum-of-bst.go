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
  val := rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
  if root.Val >= low && root.Val <= high {
    val += root.Val
  }
  return val
}