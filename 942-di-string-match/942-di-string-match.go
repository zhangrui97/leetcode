func diStringMatch(s string) []int {
  min, max := 0, 0
  ans := make([]int, len(s) + 1)
  for i := 0; i < len(s); i ++ {
    switch s[i] {
      case 'I':
        max++
        ans[i+1] = max
      case 'D':
        min--
        ans[i+1] = min
    }
  }
  if min < 0 {
    for i := 0; i <= len(s); i ++ {
      ans[i] -= min
    }
  }
  return ans
}