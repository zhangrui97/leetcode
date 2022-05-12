/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
  ans := -1000
  var helper func(t *TreeNode) int
  helper = func(t *TreeNode) int {
    if t == nil {
      return 0
    }
    l := helper(t.Left)
    r := helper(t.Right)
    sum := t.Val
    if l > 0 && r > 0 {
      if l + r + sum > ans {
        ans = sum + l + r
      }
      if l > r {
        sum += l
      } else {
        sum += r
      }
    } else if l > 0 {
      sum += l
    } else if r > 0 {
      sum += r
    }
    if sum > ans {
      ans = sum
    }
    return sum
  }
  helper(root)
  return ans
}