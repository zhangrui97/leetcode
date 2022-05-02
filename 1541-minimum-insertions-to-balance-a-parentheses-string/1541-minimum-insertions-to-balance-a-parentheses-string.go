func minInsertions(s string) int {
  res, need := 0, 0
  for _, ch := range s {
    if ch == '(' {
      need += 2
      if need%2 == 1 {
        res++
        need--
      } 
    } else {
      need--
      if need == -1 {
        res++
        need = 1
      }
    }
  }
  return res + need
}