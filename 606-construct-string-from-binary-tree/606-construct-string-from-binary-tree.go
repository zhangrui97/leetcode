/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(path *string, t *TreeNode) {
  if t == nil { return }
  *path += strconv.Itoa(t.Val)
  if t.Left == nil && t.Right == nil { return }
  *path += "("
  dfs(path, t.Left)
  *path += ")"
  if t.Right != nil {
    *path += "("
    dfs(path, t.Right)
    *path += ")"
  }
}
func tree2str(root *TreeNode) string {
  ans := ""
  dfs(&ans, root) 
  return ans
}