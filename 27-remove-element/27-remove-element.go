func removeElement(nums []int, val int) int {
  s := 0
  for _, v := range nums {
    if v != val {
      nums[s] = v
      s++
    }
  }
  return s
}