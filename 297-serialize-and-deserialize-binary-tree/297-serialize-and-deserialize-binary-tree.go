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
  data string
  pos int
}

func Constructor() Codec {
  return Codec{}
}

func (this *Codec) SerializeTree(tree *TreeNode) {
  if tree == nil {
    this.data += "#"
  } else {
    this.data += "("
    this.data += strconv.Itoa(tree.Val)
    this.data += "|"
    this.SerializeTree(tree.Left)
    this.SerializeTree(tree.Right)
    this.data += ")"
  }
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
  this.data = ""
  this.SerializeTree(root)
  return this.data
}

func (this *Codec) ParseTree() *TreeNode {
  if this.data[this.pos] == '(' {
    this.pos++
    lo := this.pos
    for ; this.data[this.pos] != '|'; this.pos++ {}
    val, _ := strconv.Atoi(string(this.data[lo:this.pos]))
    this.pos++
    tree := TreeNode{val, this.ParseTree(), this.ParseTree()}
    this.pos++
    return &tree
  } else {
    this.pos++
    return nil
  }
}
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
  this.data = data
  this.pos = 0
  return this.ParseTree()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */