/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
  ans := 0
  var isBST func(t *TreeNode)(int,int,int,bool)
  isBST = func(t *TreeNode)(int,int,int,bool) {
    if t == nil {
      return 0, 0, 0, true
    }
    minVal, maxVal, sum := t.Val, t.Val, t.Val
    if t.Left == nil && t.Right == nil {
      if sum > ans {
        ans = sum
      }
      return t.Val, t.Val, sum, true
    }
    if t.Left != nil {
      if t.Left.Val >= t.Val {
        return 0, 0, 0, false
      }
      if min, max, ls, isBst := isBST(t.Left); isBst && max < t.Val {
        minVal = min
        sum += ls
      } else {
        return 0, 0, 0, false
      }
    }
    if t.Right != nil {
      if t.Right.Val <= t.Val {
        return 0, 0, 0, false
      } 
      if min, max, rs, isBst := isBST(t.Right); isBst && min > t.Val{
        maxVal = max
        sum += rs
      } else {
        return 0, 0, 0, false
      }
    }
    if sum > ans {
      ans = sum
    }
    return minVal, maxVal, sum, true
  }
  var dfs func(t *TreeNode)
  dfs = func(t *TreeNode) {
    if _, _, _, is := isBST(t); is {
      return
    }
    dfs(t.Left)
    dfs(t.Right)
  }
  dfs(root)
  return ans
}