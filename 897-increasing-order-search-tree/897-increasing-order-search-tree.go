/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, tail *TreeNode)*TreeNode {
  if t == nil { return tail }
  l := dfs(t.Left, t)
  t.Left = nil
  t.Right = dfs(t.Right, tail)
  return l
}

func increasingBST(root *TreeNode) *TreeNode {
  return dfs(root, nil)
}