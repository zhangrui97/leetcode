/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func conn(left *Node, right *Node) {
  if left != nil {
    left.Next = right
    conn(left.Right, right.Left)
  }
}

func connect(root *Node) *Node {
  if root != nil {
    connect(root.Left)
    connect(root.Right)
    conn(root.Left, root.Right)
  }
  return root	
}