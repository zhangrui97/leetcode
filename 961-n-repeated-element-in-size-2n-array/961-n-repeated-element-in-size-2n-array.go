func repeatedNTimes(nums []int) int {
  flags := make([]bool, 10001)
  for _, v := range nums {
    if flags[v] {
      return v
    } else {
      flags[v] = true
    }
  }
  return 0
}