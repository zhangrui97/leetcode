func maxProfit(prices []int) int {
  ans := 0
  for i := 1; i < len(prices); i++ {
    if d := prices[i] - prices[i-1]; d > 0 {
      ans += d
    }
  }
  return ans
}