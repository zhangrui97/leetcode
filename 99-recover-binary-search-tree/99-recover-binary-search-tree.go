/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
  var pre, n0, n1 *TreeNode
  var dfs func(t *TreeNode)
  dfs = func(t *TreeNode) {
    if t == nil { return }
    dfs(t.Left)
    if pre == nil {
      pre = t
    }
    if pre.Val > t.Val {
      if n0 == nil {
        n0 = pre
      }
      n1 = t
    }
    pre = t
    dfs(t.Right)
  }    
  dfs(root)
  n0.Val, n1.Val = n1.Val, n0.Val
}