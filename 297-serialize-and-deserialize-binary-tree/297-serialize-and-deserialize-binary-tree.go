/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strconv"

type Codec struct {
}

func Constructor() Codec {
  return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
  builder := strings.Builder{}
  var helper func(t *TreeNode)
  helper = func(t *TreeNode) {
    if t == nil {
      builder.WriteString("#|")
    } else {
      builder.WriteString(strconv.Itoa(t.Val))
      builder.WriteByte('|')
      helper(t.Left)
      helper(t.Right)
    }
  }
  helper(root)
  return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
  items := strings.Split(data, "|")
  var parser func() *TreeNode
  parser = func() *TreeNode {
    if items[0] == "#" {
      items = items[1:]
      return nil
    } else {
      val, _ := strconv.Atoi(items[0])
      items = items[1:]
      return &TreeNode{val, parser(), parser()}
    }
  }
  return parser()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */