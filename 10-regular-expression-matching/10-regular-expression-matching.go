func isMatch(s string, p string) bool {
  ls, lp := len(s), len(p)
  if ls > 0 && lp == 0 { return false }
  if ls == 0 {
    if lp % 2 == 0 {
      for i := 1; i < lp; i += 2 {
        if p[i] != '*' { return false }
      }
      return true
    } else {
      return false
    }
  }
  if lp == 1 { 
    if ls > 1 { return false }
    return p[0] == '.' || p[0] == s[0]
  }
  if p[1] == '*' {
    return (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p) || isMatch(s, p[2:])
  } 
  
  return (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[1:])
}