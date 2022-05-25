/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
  res := [][]int{}
  var dfs func(t *TreeNode, sum int, path []int)
  dfs = func(t *TreeNode, sum int, path []int) {
    if t == nil { return }
    if t.Left == nil && t.Right == nil {
      if t.Val == sum {
        l := len(path)
        finalPath := make([]int, l + 1)
        copy(finalPath, path)
        finalPath[l] = sum
        res = append(res, finalPath)
      }
      return
    }
    nextSum := sum - t.Val
    path = append(path, t.Val)
    dfs(t.Left, nextSum, path)
    dfs(t.Right, nextSum, path)
  }
  dfs(root, targetSum, []int{})
  return res
}