/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
  l := len(preorder)  
  if l == 0 {
    return nil
  }
  val := preorder[0]
  if l == 1 {
    return &TreeNode{val, nil, nil}
  }
  valLeft := preorder[1]
  lenLeft := -1
  for i, v := range postorder {
    if v == valLeft {
      lenLeft = i
      break
    }
  }
  return &TreeNode{val,
                   constructFromPrePost(preorder[1:lenLeft+2], postorder[:lenLeft+1]),
                   constructFromPrePost(preorder[lenLeft+2:], postorder[lenLeft+1:l-1])}
}