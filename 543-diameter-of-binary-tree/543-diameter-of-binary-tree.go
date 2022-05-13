/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
  ans := 0
  var maxdepth func(t *TreeNode)int
  maxdepth = func(t *TreeNode)int {
    if t == nil {
      return -1
    }
    if t.Left == nil && t.Right == nil {
      return 0
    }
    l, r := maxdepth(t.Left), maxdepth(t.Right)
    length := l + r + 2
    if length > ans {
      ans = length
    }
    if l > r {
      return l + 1
    } else {
      return r + 1
    }
  }
  maxdepth(root)
  return ans
}