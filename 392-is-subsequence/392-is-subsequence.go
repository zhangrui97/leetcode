func isSubsequence(s string, t string) bool {
  sl, tl := len(s), len(t)
  if sl == 0 { return true }
  if tl < sl { return false }
  
  i, j := 0, 0
  for ; i < len(s); i++ {
    ch := s[i]
    for ; j < len(t); j++ {
      if t[j] == ch {
        j++
        break
      }
    }
    if j == len(t) {
      if t[j-1] != ch {
        return false
      }
      i++
      break
    }
  }
  return i == len(s)
}