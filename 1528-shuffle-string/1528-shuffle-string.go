func restoreString(s string, indices []int) string {
  buffer := make([]byte, len(s))
  for i,v := range indices {
    buffer[v] = s[i]
  }
  return string(buffer)
}