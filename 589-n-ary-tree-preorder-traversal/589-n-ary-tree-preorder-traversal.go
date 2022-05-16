/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
  ans := make([]int, 0)
  var traverse func(t *Node)
  traverse = func(t *Node) {
    if t == nil { return }
    ans = append(ans, t.Val)
    for _, c := range t.Children {
      traverse(c)
    }
  }
  traverse(root)
  return ans
}