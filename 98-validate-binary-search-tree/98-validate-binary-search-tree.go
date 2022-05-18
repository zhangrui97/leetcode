/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
  var inRange func(t *TreeNode, min, max int)bool
  inRange = func(t *TreeNode, min, max int)bool {
    if t == nil { return true } 
    if t.Left == nil && t.Right == nil {
      return t.Val >= min && t.Val <= max
    }
    if t.Left != nil {
      if t.Left.Val >= t.Val {
        return false
      }
      if result := inRange(t.Left, min, t.Val-1); !result {
        return false
      }
    } 
    if t.Right != nil {
      if t.Val >= t.Right.Val {
        return false
      } 
      if result := inRange(t.Right, t.Val+1, max); !result {
        return false
      }
    }
    return true
  }
  return inRange(root, math.MinInt32, math.MaxInt32)
}