func dist(pre, i, next int) int {
  dp, dn := i - pre, next - i
  if dp > dn {
    return dn
  } else {
    return dp
  }
}

func shortestToChar(s string, c byte) []int {
  ans := make([]int, len(s))
  pre := 0
  next := -len(s)
  j := 0
  for i := 0; i < len(s); i++ {
    if s[i] == c {
      pre = next
      if pre >= 0 {
        ans[i] = 0
        j = pre+1
      }
      next = i
      for ;j < next; j++ {
        ans[j] = dist(pre, j, next)        
      }
    }
  }
  ans[next] = 0
  for j = next+1; j < len(s); j++ {
    ans[j] = j - next
  }
  return ans
}