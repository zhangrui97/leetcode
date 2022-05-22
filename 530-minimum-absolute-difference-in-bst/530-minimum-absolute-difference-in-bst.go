/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func min(a, b int)int {
  if a > b {
    return b
  } else {
    return a
  }
}

func getMinimumDifference(root *TreeNode) int {
  result := 10000
  pre := 10001
  var dfs func(t *TreeNode)
  dfs = func(t *TreeNode) {
    if t == nil { return }
    dfs(t.Left)
    if pre == 10001 {
      pre = t.Val
    } else {
      diff := t.Val - pre
      result = min(result, diff)
      pre = t.Val
    }
    dfs(t.Right)
  }
  dfs(root)
  return result
}