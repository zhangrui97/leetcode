type Context struct {
  result []*TreeNode
  cache map[string]int
}

func (this *Context)Traverse(t *TreeNode) string {
  if t == nil {
    return "_"
  }
  
  s := fmt.Sprintf("%d|%s|%s", t.Val, this.Traverse(t.Left), this.Traverse(t.Right))
  if v, ok := this.cache[s]; ok {
    if v == 1 {
      this.result = append(this.result, t)
    }
    this.cache[s] = v + 1
  } else {
    this.cache[s] = 1
  }
  fmt.Println(s)
  return s
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
  ctx := &Context{make([]*TreeNode, 0), make(map[string]int)}
  ctx.Traverse(root)
  return ctx.result
}