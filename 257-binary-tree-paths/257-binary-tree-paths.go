/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, path string, res *[]string) {
  if t == nil { return }
  path += strconv.Itoa(t.Val)
  if t.Left == nil && t.Right == nil {
    *res = append(*res, path)
    return
  }
  path += "->"
  dfs(t.Left, path, res)
  dfs(t.Right, path, res)
}
func binaryTreePaths(root *TreeNode) []string {
  res := []string{}
  dfs(root, "", &res)
  return res
}