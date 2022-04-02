func kthDistinct(arr []string, k int) string {
  counter := make(map[string]int, len(arr))
  for _, v := range arr {
    counter[v]++
  }
  count := 1
  for _, v := range arr {
    if counter[v] == 1 {
      if count == k {
        return v
      }
      count ++
    }
  }
  return ""
}