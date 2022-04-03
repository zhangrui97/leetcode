func uniqueOccurrences(arr []int) bool {
  n := len(arr)
  counter := make(map[int]int, n)
  for _, v := range arr {
    counter[v]++
  }
  cnt2 := make(map[int]int, n)
  for _, v := range counter {
    cnt2[v]++
    if cnt2[v] > 1 {
      return false
    }
  }
  return true
}