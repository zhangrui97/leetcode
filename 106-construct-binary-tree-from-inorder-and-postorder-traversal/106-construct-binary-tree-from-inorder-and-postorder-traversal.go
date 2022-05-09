/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
  l := len(inorder)
  if l == 0 {
    return nil
  }  
  val := postorder[l - 1]
  if inorder[0] == val {
    return &TreeNode{val, nil, buildTree(inorder[1:], postorder[:l-1])}
  }
  idx := -1
  for i, v := range inorder {
    if v == val {
      idx = i
      break
    }
  }
  return &TreeNode{val, 
                   buildTree(inorder[:idx], postorder[:idx]),
                   buildTree(inorder[idx+1:], postorder[idx:l-1])}
}