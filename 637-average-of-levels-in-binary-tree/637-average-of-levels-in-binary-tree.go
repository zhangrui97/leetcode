/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
  ans := []float64{}
  current := []*TreeNode{root}
  for len(current) > 0 {
    sum := 0
    next := []*TreeNode{}
    for _, v := range current {
      sum += v.Val
      if v.Left != nil { next = append(next, v.Left) }
      if v.Right != nil { next = append(next, v.Right) }
    }  
    ans = append(ans, float64(sum) / float64(len(current)))
    current = next
  }
  return ans
}