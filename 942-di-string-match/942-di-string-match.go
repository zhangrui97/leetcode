func diStringMatch(s string) []int {
  min, max := 0, len(s)
  ans := make([]int, len(s), len(s) + 1)
  for i, v := range s {
    switch v {
      case 'I':
        ans[i] = min
        min++
      case 'D':
        ans[i] = max
        max--
    }
  }
  return append(ans, min)
}