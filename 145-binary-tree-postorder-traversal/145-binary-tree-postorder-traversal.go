/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, res *[]int) {
  if t != nil {
    dfs(t.Left, res)
    dfs(t.Right, res)
    *res = append(*res, t.Val)
  }
}

func postorderTraversal(root *TreeNode) []int {
  res := []int{}
  dfs(root, &res)
  return res
}