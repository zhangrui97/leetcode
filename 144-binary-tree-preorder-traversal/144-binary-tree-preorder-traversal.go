/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
  ans := make([]int, 0)
  var preorder func(t *TreeNode)
  preorder = func(t *TreeNode) {
    if t != nil {
      ans = append(ans, t.Val)
      preorder(t.Left)
      preorder(t.Right)
    }
  }
  preorder(root)
  return ans
}