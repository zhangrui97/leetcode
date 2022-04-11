func maxEnvelopes(envelopes [][]int) int {
  l := len(envelopes)
  if l == 1 { return 1 }
  sort.Slice(envelopes, func (i, j int) bool {
    if envelopes[i][0] < envelopes[j][0] {
      return true
    } else if envelopes[i][0] == envelopes[j][0] {
      return envelopes[i][1] > envelopes[j][1]
    } else {
      return false
    }
  })
  buffer := make([]int, 0, len(envelopes))
  for _, v := range envelopes {
    h := v[1]
    l, r := 0, len(buffer)
    for l < r {
      m := (l + r) / 2
      if buffer[m] < h {
        l = m + 1
      } else {
        r = m
      }
    }
    if l == len(buffer) {
      buffer = append(buffer, h)
    } else {
      buffer[l] = h
    }
  }
  return len(buffer)
}