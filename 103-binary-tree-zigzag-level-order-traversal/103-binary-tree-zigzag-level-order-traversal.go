/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
  ans := make([][]int, 0)
  p, q := make([]*TreeNode, 0), make([]*TreeNode, 0)
  if root != nil {
    q = append(q, root)
  }
  even := true
  for len(q) > 0 {
    level := make([]int, 0, len(q))
    for i := len(q)-1; i >= 0; i-- {
      level = append(level, q[i].Val)
      if even {
        if q[i].Left != nil {
          p = append(p, q[i].Left)
        }
        if q[i].Right != nil {
          p = append(p, q[i].Right)
        }
      } else {
        if q[i].Right != nil {
          p = append(p, q[i].Right)
        }
        if q[i].Left != nil {
          p = append(p, q[i].Left)
        }
      }
    }
    ans = append(ans, level)
    p, q = q[:0], p
    even = !even
  }
  return ans
}