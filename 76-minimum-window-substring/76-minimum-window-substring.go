func buildCounter(t string) map[byte]int {
  result := make(map[byte]int)
  for _, v := range t {
    result[byte(v)] += 1
  }
  return result
}

func minWindow(s string, t string) string {
  if len(t) > len(s) {return ""}
  counter := buildCounter(t)
  target := len(counter)
  left, length := 0, len(s) + 1
  l, r := 0, 0
  result := make(map[byte]int)
  matchNumber := 0
  for r < len(s) {
    chr := s[r]
    r++
    if cntr, ok := counter[chr]; ok {
      result[chr]++
      if result[chr] == cntr {
        matchNumber++
        for matchNumber == target {
          chl := s[l]
          if cntl, okl := counter[chl]; okl {
            result[chl]--
            if result[chl] < cntl {
              matchNumber--
              if r - l < length {
                left, length = l, r - l
              }
            }
          }
          l++
        }
      }
    }
  }
  if length > len(s) {
    return ""
  } else {
    return s[left:left+length]
  }
}