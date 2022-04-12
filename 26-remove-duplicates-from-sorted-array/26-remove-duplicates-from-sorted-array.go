func removeDuplicates(nums []int) int {
  s := 0
  for _, v := range nums {
    if v != nums[s] {
      s++
      nums[s] = v
    }
  }
  return s+1
}