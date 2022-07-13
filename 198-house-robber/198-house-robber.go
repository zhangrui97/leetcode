func max(a, b int)int { if a > b { return a } else { return b } }
func rob(nums []int) int {
  dp0, dp1 := 0, 0
  for _, n := range nums {
    dp0, dp1 = dp1, max(dp1, dp0 + n)
  }
  return dp1
}