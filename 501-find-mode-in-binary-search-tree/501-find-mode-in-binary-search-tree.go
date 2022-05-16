/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
  ans := make([]int, 0)
  var traverse func(t *TreeNode)
  pre := -100001
  maxOccured := 0
  occurred := 0
  traverse = func(t *TreeNode) {
    if t == nil {
      return
    }
    traverse(t.Left)
    if t.Val == pre {
      occurred++
    } else {
      occurred = 1
      pre = t.Val
    }
    if occurred == maxOccured {
      ans = append(ans, t.Val)  
    }
    if occurred > maxOccured {
      ans = ans[:0]
      ans = append(ans, t.Val)
      maxOccured = occurred
    }      

    traverse(t.Right)
  }
  traverse(root)
  return ans
}