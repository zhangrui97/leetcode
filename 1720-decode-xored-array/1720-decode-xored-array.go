func decode(encoded []int, first int) []int {
  pre := first
  ans := make([]int, len(encoded) + 1)
  ans[0] = first
  for i, v := range encoded {
    pre = pre ^ v
    ans[i+1] = pre
  }  
  return ans
}