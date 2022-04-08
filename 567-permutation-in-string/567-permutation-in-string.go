func checkInclusion(s1 string, s2 string) bool {
  l1, l2 := len(s1), len(s2)
  if l2 < l1 { return false }
  target := make(map[byte]int, len(s1))
  for _, v := range s1 { target[byte(v)]++ }
  targetl := len(target)
  window := make(map[byte]int, len(s1))
  valid := 0
  l, r := 0, 0
  for r < l2 {
    rc := s2[r]
    r++
    if val, ok := target[rc]; ok {
      window[rc]++
      if val == window[rc] {
        valid ++
        if (valid == targetl) {
          return true
        }
      }
    }
    if r >= l1 {
      lc := s2[l]
      if val, ok := target[lc]; ok {
        if val == window[lc] {
          valid--
        }
        window[lc]--
      }
      l++
    }
  }
  return false
}