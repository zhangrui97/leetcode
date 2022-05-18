/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
  size := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}
  var genTrees func(s, e int)[]*TreeNode 
  genTrees = func(s, e int)[]*TreeNode {
    if s >= e { 
      return []*TreeNode{nil}
    }
    if s+1 == e { 
      return []*TreeNode{&TreeNode{s, nil, nil}}
    }
    if s+2 == e {
      return []*TreeNode{&TreeNode{s, nil, &TreeNode{s+1, nil, nil}}, &TreeNode{s+1, &TreeNode{s, nil, nil}, nil}}
    }
    ans := make([]*TreeNode, 0, size[n])
    for i := s; i < e; i++ {
      for _, vl := range genTrees(s, i) {
        for _, vr := range genTrees(i+1, e) {
          ans = append(ans, &TreeNode{i, vl, vr})
        }
      }
    }
    return ans
  }
  return genTrees(1, n+1)        
}