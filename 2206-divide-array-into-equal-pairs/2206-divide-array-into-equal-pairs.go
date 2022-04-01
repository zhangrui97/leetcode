func divideArray(nums []int) bool {
  dict := make(map[int]bool, len(nums))
  result := 0
  for _, v := range nums {
    result ^= v
    dict[v] = !dict[v]
  }
  if result != 0 {
    return false
  }
  for _, v := range dict {
    if v {
      return false
    }
  }
  return true
}