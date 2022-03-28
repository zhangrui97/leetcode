func numIdenticalPairs(nums []int) int {
  l := len(nums)
  count := 0
  for i := 0; i < l - 1; i ++ {
    for j := i + 1; j < l; j++ {
      if nums[i] == nums[j] {
        count ++
      }
    }
  }
  return count
}