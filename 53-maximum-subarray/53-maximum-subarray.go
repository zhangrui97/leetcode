func maxSubArray(nums []int) int {
  dp := make([]int, len(nums))
  dp[0] = nums[0]
  ans := dp[0]
  for i := 1; i < len(nums); i++ {
    if dp[i-1] < 0 {
      dp[i] = nums[i]
    } else {
      dp[i] = nums[i] + dp[i-1]
    }
    if dp[i] > ans {
      ans = dp[i]
    }
  }
  return ans
}