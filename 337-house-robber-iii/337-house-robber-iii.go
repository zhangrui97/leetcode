/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int)int { if a > b { return a } else { return b } }

func rob(root *TreeNode) int {
  var robRec func(t *TreeNode) (int, int);
  robRec = func(t *TreeNode) (int, int) {
    if t == nil { return 0, 0 }
    if t.Left == nil && t.Right == nil {
      return t.Val, 0
    }
    l, lc := robRec(t.Left)
    r, rc := robRec(t.Right)
    return max(t.Val + lc + rc, l + r), l + r
  }
  ans, _ := robRec(root)
  return ans
}