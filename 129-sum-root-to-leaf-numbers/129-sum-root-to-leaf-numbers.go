/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
  res := 0
  var dfs func(t *TreeNode, sum int)
  dfs = func(t *TreeNode, sum int) {
    if t == nil { return }
    sum = sum * 10 + t.Val
    if t.Left == nil && t.Right == nil {
      res += sum
      return
    }
    dfs(t.Left, sum)
    dfs(t.Right, sum)
  }
  dfs(root, 0)
  return res
}