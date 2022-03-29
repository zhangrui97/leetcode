func createTargetArray(nums []int, index []int) []int {
  ans := make([]int, len(nums))
  for i, v := range index {
    if v < i {
      copy(ans[v+1:i+1], ans[v:i])
    }
    ans[v] = nums[i]
  }
  return ans
}