/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  }
  idx, val := -1, -1
  for i, v := range nums {
    if v > val {
      idx = i
      val = v
    }
  }
  return &TreeNode{val, 
                   constructMaximumBinaryTree(nums[:idx]), 
                   constructMaximumBinaryTree(nums[idx+1:])}
}