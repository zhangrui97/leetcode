/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func toArray(res *[]int, t *TreeNode) {
  if t == nil { return }
  toArray(res, t.Left)
  *res = append(*res, t.Val)
  toArray(res, t.Right)
}

func findTarget(root *TreeNode, k int) bool {
  arr := []int{}
  toArray(&arr, root)
  l, r := 0, len(arr) -1
  for l < r {
    sum := arr[l] + arr[r]
    if sum < k {
      l++
    } else if sum > k {
      r--
    } else {
      return true
    }
  }
  return false
}