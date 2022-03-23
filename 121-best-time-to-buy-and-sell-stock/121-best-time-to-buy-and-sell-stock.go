func min(a, b int) int {
  if a > b {
    return b
  } else {
    return a
  }
}
func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}
func maxProfit(prices []int) int {
  dp := math.MaxInt32
  ans := 0
  for _, v := range prices {
    dp = min(dp, v)
    ans = max(ans, v - dp)
  }
  return ans
}