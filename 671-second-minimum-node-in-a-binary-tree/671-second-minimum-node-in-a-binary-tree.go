/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, min int, ans *int) {
  if t == nil { return }
  if t.Val == min {
    dfs(t.Left, min, ans)
    dfs(t.Right, min, ans)
    return
  }
  if t.Val > min && (*ans == -1 || t.Val < *ans) { 
    *ans = t.Val 
  }
}
func findSecondMinimumValue(root *TreeNode) int {
  if root.Left == nil && root.Right == nil { return -1 }
  val := root.Left.Val + root.Right.Val - root.Val
  if val == root.Val + 1 { return val }
  ans := val
  if ans == root.Val {
    ans = -1
  }
  if root.Left.Val == val {
    dfs(root.Right, root.Val, &ans)
  } else {
    dfs(root.Left, root.Val, &ans)
  }
  return ans
}