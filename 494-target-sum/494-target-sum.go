func findTargetSumWays(nums []int, target int) int {
  n := len(nums)
  ans := 0
  var trackback func(i, t int)
  trackback = func(i, t int) {
    if i == n {
      if t == 0 { ans++ }
      return
    }
    trackback(i+1, t + nums[i])
    trackback(i+1, t - nums[i])
  }
  trackback(0, target)
  return ans
}