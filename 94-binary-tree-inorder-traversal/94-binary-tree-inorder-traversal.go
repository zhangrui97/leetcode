/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
  ans := make([]int, 0)
  var traverse func(t *TreeNode)
  traverse = func(t *TreeNode) {
    if t == nil {
      return
    }
    traverse(t.Left)
    ans = append(ans, t.Val)
    traverse(t.Right)
  }
  traverse(root)
  return ans
}