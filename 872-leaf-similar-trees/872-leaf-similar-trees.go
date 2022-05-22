/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafs(root *TreeNode)[]int {
  res := []int{}
  var dfs func(t *TreeNode)
  dfs = func(t *TreeNode) {
    if t == nil { return }
    if t.Left == nil && t.Right == nil {
      res = append(res, t.Val)
      return
    }
    dfs(t.Left)
    dfs(t.Right)
  }
  dfs(root)
  return res
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
  res1, res2 := leafs(root1), leafs(root2)
  if len(res1) != len(res2) { return false }
  for i := 0; i < len(res1); i++ {
    if res1[i] != res2[i] {
      return false
    }
  }
  return true
}