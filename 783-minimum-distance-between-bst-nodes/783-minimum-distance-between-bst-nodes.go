/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
  result := 10000
  pre := -1
  var dfs func(t *TreeNode)
  dfs = func(t *TreeNode) {
    if t == nil { return }
    dfs(t.Left)
    if pre == -1 {
      pre = t.Val
    } else {
      diff := t.Val - pre
      if diff < result {
        result = diff
      }
      pre = t.Val
    }
    dfs(t.Right)
  }
  dfs(root)
  return result
}