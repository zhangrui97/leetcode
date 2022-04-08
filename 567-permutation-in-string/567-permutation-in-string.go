func checkInclusion(s1 string, s2 string) bool {
  target := make(map[byte]int, len(s1))
  for _, v := range s1 { target[byte(v)]++ }
  targetl := len(target)
  window := make(map[byte]int, len(s1))
  valid := 0
  l, r := 0, 0
  for r < len(s2) {
    rc := s2[r]
    r++
    if val, ok := target[rc]; ok {
      window[rc]++
      if val == window[rc] {
        valid ++
      }
    }
    for valid == targetl {
      lc := s2[l]
      if val, ok := target[lc]; ok {
        if val == window[lc] {
          if r - l == len(s1) {
            return true
          }
          valid--
        }
        window[lc]--
      }
      l++
    }
  }
  return false
}