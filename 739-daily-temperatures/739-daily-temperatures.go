func dailyTemperatures(temperatures []int) []int {
  ans := make([]int, len(temperatures))
  stack := make([]int, 0, len(temperatures))
  for i, v := range temperatures {
    for l:=len(stack); l > 0 && v > temperatures[stack[l-1]]; l-- {
      j := stack[l-1]
      ans[j] = i - j
      stack = stack[:l-1]
    }
    stack = append(stack, i)
  }
  return ans
}