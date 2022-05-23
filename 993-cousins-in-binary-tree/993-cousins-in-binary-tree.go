/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
  depth1, depth2 := -1, -1
  var parent1, parent2 *TreeNode
  var target int
  var dfs func(p, t *TreeNode, d int)bool
  dfs = func(p, t *TreeNode, d int)bool {
    if t == nil { return false }
    if target == 0 {
      if t.Val == x || t.Val == y {
        parent1 = p
        depth1 = d
        target = x + y - t.Val
      }
    } else {
      if t.Val == target {
        parent2 = p
        depth2 = d
        return true
      }
    }
    return dfs(t, t.Left, d + 1) || dfs(t, t.Right, d + 1)
  }
  dfs(nil, root, 0)
  return depth1 == depth2 && parent1 != parent2
}