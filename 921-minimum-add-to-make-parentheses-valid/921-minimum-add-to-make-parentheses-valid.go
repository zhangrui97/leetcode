func minAddToMakeValid(s string) int {
  l := 0
  r := 0
  for _, ch := range s {
    switch ch {
      case '(': l++
      case ')': 
        if l > 0 { 
          l-- 
        } else {
          r++
        } 
    }
  }
  return r + l
}