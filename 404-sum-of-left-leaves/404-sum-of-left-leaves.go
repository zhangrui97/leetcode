/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, isLeft bool)int {
  if t == nil { return 0 }
  if t.Left == nil && t.Right == nil {
    if isLeft {
      return t.Val
    } else {
      return 0
    }
  }
  return dfs(t.Left, true) + dfs(t.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
  return dfs(root.Left, true) + dfs(root.Right, false)
}