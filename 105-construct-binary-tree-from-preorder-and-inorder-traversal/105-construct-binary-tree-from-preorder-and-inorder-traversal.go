/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(inorder) == 0 {
    return nil
  }
  val := preorder[0]
  if val == inorder[0] {
    return &TreeNode{val, nil, buildTree(preorder[1:], inorder[1:])}
  }
  idxVal := -1
  for i, v := range inorder {
    if val == v {
      idxVal = i
      break
    }
  }
  return &TreeNode{val, 
                   buildTree(preorder[1:idxVal+1], inorder[:idxVal]), 
                   buildTree(preorder[idxVal+1:], inorder[idxVal+1:])}
}