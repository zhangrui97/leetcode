/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
  if root == nil { return nil }
  ans := make([]int, 0)
  for _, c := range root.Children {
    ans = append(ans, postorder(c)...)
  }
  ans = append(ans, root.Val)
  return ans
}