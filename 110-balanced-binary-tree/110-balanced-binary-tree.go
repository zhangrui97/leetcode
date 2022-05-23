/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}
func dfs(root *TreeNode, d int)(bool, int) {
  if root == nil { return true, d }
  if root.Left == nil && root.Right == nil { return true, d + 1 }
  if root.Left == nil {
    if root.Right.Left == nil && root.Right.Right == nil {
      return true, d + 2
    } else {
      return false, d + 3
    }
  }
  if root.Right == nil {
    if root.Left.Left == nil && root.Left.Right == nil {
      return true, d + 2
    } else {
      return false, d + 3
    }
  }
  l, dl := dfs(root.Left, d+1)
  if !l { return false, dl }
  r, dr := dfs(root.Right, d+1)  
  if !r { return false, dr }
  if dr - dl > 1 || dl - dr > 1 {
    return false, dl
  } else {
    return true, max(dr, dl)
  }
}
func isBalanced(root *TreeNode) bool {
  res, _ := dfs(root, 0)
  return res
}