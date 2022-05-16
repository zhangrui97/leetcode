/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func maxDepth(root *Node) int {
  if root == nil {
    return 0
  } else if len(root.Children) == 0 {
    return 1
  }
  max := 0
  for _, c := range root.Children {
    d := maxDepth(c)
    if d > max {
      max = d
    }
  }
  return max + 1
}