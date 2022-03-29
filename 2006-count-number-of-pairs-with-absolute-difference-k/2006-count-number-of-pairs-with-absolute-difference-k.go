func countKDifference(nums []int, k int) int {
  if len(nums) <= 1 {
    return 0
  }
  ans := 0
  gt, le := nums[0] + k, nums[0] - k
  for i := 1; i < len(nums); i++ {
    if nums[i] == gt || nums[i] == le {
      ans ++
    }
  }
  
  return ans + countKDifference(nums[1:], k)
}