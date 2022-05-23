/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, sum, target int) bool {
  if t == nil { return false }
  sum += t.Val
  if t.Left == nil && t.Right == nil {
    if sum == target {
      return true
    }
  }  
  return dfs(t.Left, sum, target) || dfs(t.Right, sum, target)
}
func hasPathSum(root *TreeNode, targetSum int) bool {
  return dfs(root, 0, targetSum)  
}