/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int)int {
  if a > b { return a }
  return b
}

func level(t *TreeNode)int {
  if t == nil { return 0 }
  return max(level(t.Left), level(t.Right)) + 1
}

func deepestSum(t *TreeNode, d int, sum *int) {
  if t == nil { return }
  if d == 1 {
    *sum += t.Val
  }
  deepestSum(t.Left, d-1, sum)
  deepestSum(t.Right, d-1, sum)
}
func deepestLeavesSum(root *TreeNode) int {
  sum := 0
  deepestSum(root, level(root), &sum)
  return sum
}