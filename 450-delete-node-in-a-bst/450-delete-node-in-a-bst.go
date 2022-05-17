/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
  if root == nil { return root }
  if root.Val > key {
    root.Left = deleteNode(root.Left, key)
  }
  if root.Val == key {
    if root.Left == nil && root.Right == nil {
      return nil
    }
    if root.Left == nil {
      return root.Right
    }
    if root.Right == nil {
      return root.Left
    }
    nxt := root.Right
    for nxt.Left != nil {
      nxt = nxt.Left
    }
    root.Right = deleteNode(root.Right, nxt.Val)
    nxt.Left = root.Left
    nxt.Right = root.Right
    return nxt
  }
  root.Right = deleteNode(root.Right, key)
  return root
}