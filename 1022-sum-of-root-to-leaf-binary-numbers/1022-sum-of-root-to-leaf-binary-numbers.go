/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, v int)int {
  if t == nil { return 0 }
  v = v << 1 + t.Val
  if t.Left == nil && t.Right == nil {
    return v
  }
  return dfs(t.Left, v) + dfs(t.Right, v)
}
func sumRootToLeaf(root *TreeNode) int {
  return dfs(root, 0)
}