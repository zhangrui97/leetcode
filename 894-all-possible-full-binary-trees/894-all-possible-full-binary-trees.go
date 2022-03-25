/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
  if n % 2 == 0 { return []*TreeNode{} }
  level := (n + 1) / 2 
  dp := make([][]*TreeNode, level)
  dp[0] = []*TreeNode{&TreeNode{0, nil, nil}}
  for i := 1; i < level; i ++ {
    var si []*TreeNode
    for l := 0; l < i; l ++ {
      for _, lNode := range dp[l] {
        for _, rNode := range dp[i-l-1] {
          si = append(si, &TreeNode{0, lNode, rNode})
        }
      }
    }
    dp[i] = si
  }
  return dp[level-1]}