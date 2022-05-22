/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sum(t *TreeNode) int {
  if t == nil { return 0 }
  return t.Val + sum(t.Left) + sum(t.Right)
}

func getTilt(t *TreeNode) int {
  if t == nil { return 0 }
  l, r := sum(t.Left), sum(t.Right)
  if l > r {
    return l - r
  } else {
    return r - l
  }
}

func dfs(t *TreeNode) {
  if t == nil {
    return
  }  
  t.Val = getTilt(t)
  dfs(t.Left)
  dfs(t.Right)
}

func findTilt(root *TreeNode) int {
  dfs(root)
  return sum(root)  
}