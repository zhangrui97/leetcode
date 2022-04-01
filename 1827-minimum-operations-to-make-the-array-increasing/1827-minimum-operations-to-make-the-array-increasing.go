func minOperations(nums []int) int {
  ans := 0
  pre := nums[0] + 1
  for i := 1; i < len(nums); i ++ {
    if nums[i] < pre {
      ans += pre - nums[i]
      pre ++
    } else {
      pre = nums[i] + 1
    }
  }
  return ans
}