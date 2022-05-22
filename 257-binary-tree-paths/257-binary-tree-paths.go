/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(t *TreeNode, path []int, res *[][]int) {
  if t == nil { return }
  path = append(path, t.Val)
  if t.Left == nil && t.Right == nil {
    p := make([]int, len(path))
    copy(p, path)
    *res = append(*res, p)
    return
  }
  dfs(t.Left, path, res)
  dfs(t.Right, path, res)
}
func binaryTreePaths(root *TreeNode) []string {
  res := [][]int{}
  dfs(root, []int{}, &res)
  ans := make([]string, len(res))
  for i, v := range res {
    ans[i] = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), "->"), "[]")
  }
  return ans
}