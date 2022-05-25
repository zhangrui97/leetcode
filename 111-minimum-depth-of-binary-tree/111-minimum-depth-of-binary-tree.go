/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
  if root == nil { return 0 }
  res := 0
  level := []*TreeNode{root}
  for l := len(level); l > 0; l = len(level) {
    for i := 0; i < l; i ++ {
      t := level[0]
      level = level[1:]
      if t.Left == nil && t.Right == nil {
        return res + 1
      }
      if t.Left != nil { level = append(level, t.Left) }
      if t.Right != nil { level = append(level, t.Right) }
    }
    res++
  }  
  return res
}