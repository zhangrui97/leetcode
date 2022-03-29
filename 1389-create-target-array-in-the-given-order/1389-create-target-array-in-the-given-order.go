func createTargetArray(nums []int, index []int) []int {
  ans := make([]int, len(nums))
  for i, v := range index {
    for j := i; j > v; j-- {
      ans[j] = ans[j-1]
    }
    ans[v] = nums[i]
  }
  return ans
}