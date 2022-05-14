/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
  ans := make([][]int, 0)
  p, q := make([]*TreeNode, 0), make([]*TreeNode, 0)
  if root != nil {
    q = append(q, root)
  }
  for len(q) > 0 {
    level := make([]int, len(q))
    for i, v := range q {
      level[i] = v.Val
      if v.Left != nil {
        p = append(p, v.Left)
      }
      if v.Right != nil {
        p = append(p, v.Right)
      }
    }
    ans = append(ans, level)
    p, q = q[:0], p
  }
  return ans
}