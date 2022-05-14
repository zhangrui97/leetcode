/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
  if root == nil { 
    return 0
  }  
  counter := 0
  depth := 0
  var find func(t *TreeNode, d int)bool
  find = func(t *TreeNode, d int)bool {
    if t.Right == nil {
      if t.Left != nil {
        depth = d + 1
        counter ++
        return true
      }
      if depth == 0 {
        depth = d
        counter += 2
        return false
      } else if d > depth {
        depth = d
        return true
      }
      counter += 2
      return false
    }
    if find(t.Right, d+1) {
      return true
    }
    return find(t.Left, d+1)
  }
  if find(root, 1) {
    return (1 << depth) - 1 - counter
  } else {
    return (1 << depth) - 1
  }
}